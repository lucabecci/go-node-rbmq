const {Router} = require('express')
const RouterController = require('../controllers/index.controller')
class IndexRouter{
    _router
    _routerController
    constructor(){
        this._router = Router()
        this._routerController = new RouterController

        this.initRoutes()
    }

    initRoutes(){
        this._router.get('/', this._routerController.getIndex)
        this._router.post('/', this._routerController.newData)
    }
}

module.exports = IndexRouter