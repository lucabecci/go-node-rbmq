class Broker {
    static channel 
    static goSendCreated = false
    static loadChannel(channel){
        this.channel = channel
    }
    static async createQueue(queue){
        if(this.goSendCreated === false){
            this.channel.assertQueue(queue)
            console.log('creating queue:', queue)
            this.goSendCreated = true;
        }
    }
    static async sendMessage(queue, message){
        await this.channel.sendToQueue(queue, Buffer.from(JSON.stringify(message)))
    }
}

module.exports = Broker