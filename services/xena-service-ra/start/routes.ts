/*
|--------------------------------------------------------------------------
| Routes
|--------------------------------------------------------------------------
|
| This file is dedicated for defining HTTP routes. A single file is enough
| for majority of projects, however you can define routes in different
| files and just make sure to import them inside this file. For example
|
| Define routes in following two files
| ├── start/routes/cart.ts
| ├── start/routes/customer.ts
|
| and then import them inside `start/routes.ts` as follows
|
| import './routes/cart'
| import './routes/customer'
|
*/

import Route from '@ioc:Adonis/Core/Route'
import Env from '@ioc:Adonis/Core/Env'

Route.group(() => {

  Route.group(() => {
    Route.post('/', 'ReadersController.get')
  }).prefix('readers')

  Route.group(() => {
    Route.get('/', 'MachinesController.details')
  }).prefix('machines')

  Route.group(() => {
    Route.post('/', 'DatabasesController.injection')
  }).prefix('databases')

  Route.group(() => {
    Route.post('/subdomain-bruteforce', 'SubdomainsController.bruteForce')
    Route.post('/sublist3r', 'SubdomainsController.sublist3r')
    Route.post('/nmap', 'MappersController.nmap')
    Route.post('/websearch', 'ReadersController.webSearch')
  }).prefix('recon')

  Route.group(() => {
    Route.post('/sql-injection', 'SqlInjectionsController.scan')
    Route.post('/web-fuzzer', 'WebFuzzersController.scan')
  }).prefix('scans')

}).prefix(`${Env.get('DIR_BUSTER') ?? ''}/v1`).middleware(['auth'])
