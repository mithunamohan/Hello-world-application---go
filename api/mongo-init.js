db = db.getSiblingDB('deb'); //create db named deb

//create greeting collection
db.greetings.insertMany([
    {'greetings':'Hello World!'},
    {'greetings':'Bonjour Monde!'},
    {'greetings':'Hallo Welt!'},
]);