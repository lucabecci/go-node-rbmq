const amqp = require('amqplib')

class BrokerFactory{
    _amqp
    _channel
    constructor(user, password){
        this._amqp = amqp
        this._channel = this.configuration(user, password)
    }
    async configuration(user, password){
        let cnn;
        let retries = 5;
        while(retries){
            try{
                cnn = await this._amqp.connect(`amqp://${user}:${password}@localhost:5672/`)
                console.log('amqp is connected')
                break
            }
            catch(e){
                retries -= 1;
                console.log("retries:", retries);
                await new Promise((res) => setTimeout(res, 5000));
            }
        }
        const channel = await cnn.createChannel();
        return channel
        
    }

    async getChannel(){
        return this._channel
    }
}

module.exports = BrokerFactory