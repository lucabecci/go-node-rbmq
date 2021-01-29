const express = require('express');
const morgan = require('morgan')
require('dotenv/config')
const BrokerFactory = require('./broker/BrokerFactory')
const Broker = require('./broker/Broker')

const IndexRouter = require('./routes/index.routes')
class Server {
    _app 
    _indexRouter
    constructor(){
        this._app = express()
        this._indexRouter = new IndexRouter

        this.initBroker()
        this.initConfig()
        this.initRoutes()
    }

    async initBroker(){
        const brokerFactory = new BrokerFactory();
        const channel = await brokerFactory.getChannel()
        Broker.loadChannel(channel)
    }

    initConfig(){
        this._app.use(express.json())
        this._app.use(express.urlencoded({extended: false}))
        this._app.use(morgan('dev'))
    }

    initRoutes(){
        this._app.use(this._indexRouter._router)
    }

    run(){
        this._app.listen(5000, () => {
            console.log("Srv on port:", 5000)
        })
    }
}


const srv = new Server

srv.run()