from os import system
from subprocess import check_output
from time import time, time_ns

class Main:
  private_key = ''
  public_key = ''
  postgres_pass = ''
  app_key = ''

  def __init__ (self):
    self.load_keys()
    print()
    print(self.public_key)
    print()
    self.postgres_pass = str(time()) + str(time_ns())
    self.app_key = str(time()) + str(time_ns())
    self.init_xena_docker_network()
    self.setup_atila_postgres_container()
    self.setup_pyramid_postgres_container()
    self.setup_xena_service_face_container()
    self.setup_xena_service_atila_container()
    self.setup_xena_service_pyramid_container()
  
  def load_keys (self):
    with open('xena.private.key', 'r') as f:
      self.private_key = f.read().replace('\n', '\\n')
      self.private_key = self.private_key[:len(self.private_key)-2]
    with open('xena.public.key', 'r') as f:
      self.public_key = f.read().replace('\n', '\\n')
      self.public_key = self.public_key[:len(self.public_key)-2]

  def init_xena_docker_network (self):
    system('docker network create xena')

  def setup_pyramid_postgres_container (self):
    system('''docker run -d --name xena-pyramid-postgres --net xena -e POSTGRES_DB=xena-pyramid -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=''' + self.postgres_pass + ''' postgres''')
    
  def setup_atila_postgres_container (self):
    system('''docker run -d --name xena-atila-postgres --net xena -e POSTGRES_DB=xena-atila -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=''' + self.postgres_pass + ''' postgres''')

  def setup_xena_service_face_container (self):
    system('''cd user-interfaces/xena-service-face && docker build -t xena-service-face . && docker run -d -p 3000:3000 --net xena --name='xena-face' xena-service-face''')

  def setup_xena_service_pyramid_container (self):
    # Regex thing is because python recognizes {{ }} as string formating.
    postgres_container_address = check_output(
      'docker inspect -f \'[[range.NetworkSettings.Networks]][[.IPAddress]][[end]]\' xena-pyramid-postgres'.replace('[[', '{' + '{').replace(']]', '}' + '}'),
      shell = True
    ).decode('utf-8').replace('\n', '')
    system('''cd services/xena-service-pyramid && docker build -t xena-service-pyramid . && docker run -d --net xena --name='xena-pyramid' -e XENA_GIT_BRANCH="stage" -e PG_HOST="''' + postgres_container_address + '''" -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD="''' + self.postgres_pass + '''" -e APP_KEY="''' + self.app_key + '''" -e TRUSTED_PUBLIC_KEY="''' + self.public_key + '''" -p 60667:60667 xena-service-pyramid''')
    print()
    print('Enter y/Y to confirm!')
    print()
    system('docker exec -ti xena-pyramid sh -c "node build/ace migration:run"')

  def setup_xena_service_atila_container (self):
    # Regex thing is because python recognizes {{ }} as string formating.
    postgres_container_address = check_output(
      'docker inspect -f \'[[range.NetworkSettings.Networks]][[.IPAddress]][[end]]\' xena-atila-postgres'.replace('[[', '{' + '{').replace(']]', '}' + '}'),
      shell = True
    ).decode('utf-8').replace('\n', '')
    system('''cd services/xena-service-atila && docker build -t xena-service-atila . && docker run -d --net xena --name='xena-atila' -e PG_HOST="''' + postgres_container_address + '''" -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD="''' + self.postgres_pass + '''" -e APP_KEY="''' + self.app_key + '''" -e TRUSTED_PUBLIC_KEY="''' + self.public_key + '''" -p 60666:60666 xena-service-atila''')
    print()
    print('Enter y/Y to confirm!')
    print()
    system('docker exec -ti xena-atila sh -c "node build/ace migration:run"')

Main()