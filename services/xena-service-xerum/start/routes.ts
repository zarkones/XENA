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
import Topic from 'App/Models/Topic'
import { v4 } from 'uuid'

Route.group(() => {

  Route.group(() => {
    Route.get('/', 'AuthorsController.get')
    Route.post('/', 'AuthorsController.insert')
  }).prefix('authors')

  Route.group(() => {
    Route.get('/', 'CommentsController.getMultiple')
    Route.post('/', 'CommentsController.insert').middleware(['auth'])
  }).prefix('comments')

  Route.group(() => {
    Route.get('/:id', 'PostsController.get')
    Route.get('/', 'PostsController.getMultiple')
    Route.post('/', 'PostsController.insert').middleware(['auth'])
  }).prefix('posts')

  Route.group(() => {
    Route.get('/', 'TopicsController.getMultiple')
  }).prefix('topics')

}).prefix('v1')

/*
Route.get('', async () => {
  await Topic.create({
    id: v4(),
    title: 'General',
    description: 'Talk about planets, humans, space, cyberspace, and other!',
  })

  await Topic.create({
    id: v4(),
    title: 'Bot Clients',
    description: 'Discussion about the bot clients, their functionality, concerns, etc...',
  })

  await Topic.create({
    id: v4(),
    title: 'Back-End Infrastructure',
    description: 'Everything related to back-end services and APIs which are meant to support the bot clients and users.',
  })

  await Topic.create({
    id: v4(),
    title: 'Front-End Interfaces',
    description: 'Everything about the user interfaces which are meant to make our life easier when interacting with services and bots.',
  })

  await Topic.create({
    id: v4(),
    title: 'Recon',
    description: 'Talk about information gathering and profiling.',
  })
})
*/