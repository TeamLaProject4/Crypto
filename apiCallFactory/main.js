const axios = require('axios');


async function main() {
    await login()
    for (let i = 0; i < 50; i++) {
        createTransaction()
    }
}


const url = "http://192.168.178.35:65524"



async function login() {
    await axios.post(url + "/frontend/confirmMnemonic", {
        mnemonic: "tenant ostrich nation lift screen inside whisper replace foam correct tree cool little announce correct excess slogan term actor crystal scout innocent viable fix",
    })
        .then((response) => {
            //do something awesome that makes the world a better place
            console.log("login succes")
        });
}

function createTransaction() {
    // Making post request
    axios.post(url + "/frontend/transaction", {
        transactionType: "stake",
        amount: "20",
        recieverPublicKey: "041294b644758b8d260741290b78aadce6a41463a7ef69438b85b2f74060baa5a6eeda1fc3144cb6b5dc7f4868beb223e79f890302aac330666466b395d1527b81"
    })
        .then((response) => {
            //do something awesome that makes the world a better place
            console.log("create trans succes")
        }).catch(e => console.log(e.response.data));
}

main()