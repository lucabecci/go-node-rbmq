const Broker = require("../broker/Broker")

class IndexRouter{
    getIndex(_req, res){
        res.send('Microservices test made with Node, Golang and RabbitMQ')
    }

    newData(req, res){
        const {title, description, author} = req.body
        const message = {
            title,
            description,
            author
        }
        try{
            Broker.createQueue('network')
            Broker.sendMessage('network', message)
            res.status(200).json({
                ok: true,
                message
            })
        }
        catch(e){
            res.status(500).json({
                ok: false,
                message: 'error to send the information'
            })
        }
    }
}

module.exports = IndexRouter