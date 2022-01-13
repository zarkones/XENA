from os import system
from subprocess import check_output

class Main:
  private_key = ''
  public_key = ''
  postgres_pass = ''
  atila_app_key = ''

  def __init__ (self):
    self.load_keys()
    print()
    print(self.public_key)
    print()
    self.postgres_pass = str(input('Enter a password for Postgres database (no need to remember): '))
    self.atila_app_key = str(input('Enter an app key for Atila service (no need to remember): '))
    self.init_xena_docker_network()
    self.setup_postgres_container()
    self.setup_xena_service_face_container()
    self.setup_xena_service_atila_container()
  
  def load_keys (self):
    with open('xena.private.key', 'r') as f:
      self.private_key = f.read().replace('\n', '\\n')
      self.private_key = self.private_key[:len(self.private_key)-2]
    with open('xena.public.key', 'r') as f:
      self.public_key = f.read().replace('\n', '\\n')
      self.public_key = self.public_key[:len(self.public_key)-2]

  def init_xena_docker_network (self):
    system('docker network create xena')
    
  def setup_postgres_container (self):
    system('''docker run -d --name xena-postgres --net xena -e POSTGRES_DB=xena-atila -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=''' + self.postgres_pass + ''' postgres''')

  def setup_xena_service_face_container (self):
    system('''cd user-interfaces/xena-service-face && docker build -t xena-service-face . && docker run -d -p 3000:3000 --net xena --name='xena-face' xena-service-face''')

  def setup_xena_service_atila_container (self):
    # Regex thing is because python recognizes {{ }} as string formating.
    postgres_container_address = check_output(
      'docker inspect -f \'[[range.NetworkSettings.Networks]][[.IPAddress]][[end]]\' xena-postgres'.replace('[[', '{' + '{').replace(']]', '}' + '}'),
      shell = True
    ).decode('utf-8').replace('\n', '')
    system('''cd services/xena-service-atila && docker build -t xena-service-atila . && docker run -d --net xena --name='xena-atila' -e PG_HOST="''' + postgres_container_address + '''" -e CORS_POLICY_ALLOWED_ORIGINS='http://127.0.0.1:3000' -e PG_PASSWORD="''' + self.postgres_pass + '''" -e APP_KEY="''' + self.atila_app_key + '''" -e TRUSTED_PUBLIC_KEY="''' + self.public_key + '''" -p 60666:60666 xena-service-atila''')
    print()
    print('Enter y/Y to confirm.')
    system('docker exec -ti xena-atila sh -c "node build/ace migration:run"')

Main()