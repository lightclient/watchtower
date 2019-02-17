const accountSid = process.env.TWILIO_SID; 
const authToken = process.env.TWILIO_AUTH_TOKEN; 
const client = require('twilio')(accountSid, authToken); 
 
function send(message, to) {
  return client.messages.create({ 
    body: message, 
    from: process.env.TWILIO_FROM_NUMBER,       
    to: to 
  })
}

module.exports = { send }
