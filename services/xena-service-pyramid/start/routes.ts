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

Route.group(() => {
  
  Route.group(() => {
    Route.get('/', 'BuildProfilesController.getMultiple')
    Route.post('/', 'BuildProfilesController.insert')
  }).prefix('build-profiles')

  Route.group(() => {
    Route.get('/list', 'BuildsController.getMultiple')
  }).prefix('builds')
  
}).prefix('v1')
  .middleware(['auth'])

// This is a state changin route, but we're assigning GET method,
// For reasons of pure convenience.
Route.get('/v1/builds', 'BuildsController.insert')