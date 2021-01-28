const amqp = require('amqplib')

class BrokerFactory{
    _amqp
    _channel
    constructor(user, password){
        this._amqp = amqp
        this._channel = this.configuration(user, password)
    }
    async configuration(user, password){
        const cnn = await this._amqp.connect(`amqp://${user}:${password}@localhost:5672/`)
        const channel = await cnn.createChannel();
        return channel
    }

    async getChannel(){
        return this._channel
    }
}

module.exports = BrokerFactory