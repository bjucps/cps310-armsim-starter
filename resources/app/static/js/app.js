const about = (html) => {
    const c = document.createElement("div");
    c.innerHTML = html;
    asticode.modaler.setContent(c);
    asticode.modaler.show();
};

const loadFile = (filename) => {
    asticode.loader.show(); // show please wait dialog

    // convert the message to send into json then send it
    astilectron.sendMessage({name: "load.file", payload: JSON.stringify({filename: filename})}, (message) => {
        asticode.loader.hide(); // hide the dialog
        // we got a response back and need to do something with it
        try {
            // display the checksum
            $('#checksum').text(message.payload.checksum);
        } catch (e) {
            // an error occurred, show a dialog with the error message
            let f = asticode.modaler.newForm();
            f.addError();
            f.addTitle("Error: " + e.message);
            asticode.modaler.setContent(f);
            asticode.modaler.show();
        }
    });
};

$(document).ready(() => {
    asticode.loader.init();
    asticode.modaler.init();
    asticode.notifier.init();

    $('#load-file').click(() =>{
        const filename = $('#file').val();
        loadFile(filename);
    });
});

document.addEventListener('astilectron-ready', function() {
    astilectron.onMessage(function(message) {
        switch (message.name) {
            case "about":
                about(message.payload);
                return {payload: "payload"};
                break;
        }
    });
});