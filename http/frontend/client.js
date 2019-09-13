require('google-protobuf')
var $;
$ = require('jquery');

const {Record,
    SubscribeRequest, SubscribeResponse,
    GetLinksRequest, GetLinksResponse,
    PutLinkRequest, PutLinkResponse,
    RemoveLinkRequest, RemoveLinkResponse} = require('./drop_pb.js');
const {DropApiClient} = require('./drop_grpc_web_pb.js');

var grpcEndpoint = "https://farnasirim.ir:18082";
if (process.env.NODE_ENV != "production") {
    grpcEndpoint = "http://localhost:8081";
}

console.log(process.env.NODE_ENV);
var client = new DropApiClient(grpcEndpoint, null, null);

var request = new SubscribeRequest();
var linksStream = client.subscribe(request, {});

window.onload = function () {
    document.getElementById('linkHref').onpaste = function(e) {
        var pastedText = undefined;
        if (window.clipboardData && window.clipboardData.getData) { // IE
                pastedText = window.clipboardData.getData('Text');
        } else if (e.clipboardData && e.clipboardData.getData) {
                pastedText = e.clipboardData.getData('text/plain');
        }
        return true;
    };

    document.getElementById("linkHref").onblur = function(e) {
        document.getElementById("linkText").value = document.getElementById("linkHref").value;
    }

    document.getElementById("linkText").onfocus = function(e) {
        document.getElementById("linkText").select();
    }

    document.getElementById("submitLink").onclick = function(e) {
        var request = new PutLinkRequest();
        var link = new Record();
        var href = document.getElementById("linkHref").value;
        var text = document.getElementById("linkText").value;
        link.setLinkaddress(href);
        link.setLinktext(text);
        request.setLink(link);

        client.putLink(request, {}, (err, response) => {
            let record = response.getLink();
            // tryCreateLink(record.getId(), record.getLinktext(), record.getLinkaddress());
        });
    }
}

elementIdFromRecordId = function(id) {
    return "record-" + id;
}

tryDeleteLink = function(id) {
    elemId = elementIdFromRecordId(id);
    elem = document.getElementById(elemId)
    if(elem) {
        elem.parentNode.removeChild(elem);
    }
}

tryCreateLink = function(id, text, addr) {
    let linksList = document.getElementById("linksList");
    const anchorTag = document.createElement("a");
    anchorTag.text = text;
    anchorTag.href = addr;

    const deleteTag = document.createElement("a");
    deleteTag.text = "[X]";
    deleteTag.href = "javascript:";
    deleteTag.onclick = function(e) {
        var request = new RemoveLinkRequest();
        var record = new Record();
        record.setId(id);
        request.setLink(record);

        client.removeLink(request, {}, (err, response) => {
            // tryDeleteLink(response.getId());
        });
    };

    const space = document.createElement("span")
    space.text = " ";

    const listElement = document.createElement("li");
    listElement.className = "linkRecord"
    listElement.dataset.sortOrder = -id;
    elemId = elementIdFromRecordId(id);
    if(document.getElementById(elemId)) {

    }
    listElement.id = elemId;

    listElement.appendChild(deleteTag);
    listElement.appendChild(space);
    listElement.appendChild(anchorTag);

    linksList.prepend(listElement);

    var $wrapper = $('.linksList');

    $wrapper.find('.linkRecord').sort(function(a, b) {
        key = +a.dataset.sortOrder - +b.dataset.sortOrder;
        return key
    })
    .appendTo($wrapper);
}

linksStream.on('data', (response) => {
    if(response.getAction() == SubscribeResponse.Action.CREATE) {
        let record = response.getRecord();
        tryCreateLink(record.getId(), record.getLinktext(), record.getLinkaddress());
    } else if(response.getAction() == SubscribeResponse.Action.REMOVE) {
        let record = response.getRecord();
        tryDeleteLink(record.getId());
    } else {
        console.log(response)
    }
});
    
// // server streaming call
// var streamRequest = new RepeatHelloRequest();
// streamRequest.setName('World');
// streamRequest.setCount(5);
// 
// var stream = client.sayRepeatHello(streamRequest, {});
// stream.on('data', (response) => {
//   console.log(response.getMessage());
// });
//   
// 
// // deadline exceeded
// var deadline = new Date();
// deadline.setSeconds(deadline.getSeconds() + 1);
// 
// client.sayHelloAfterDelay(request, {deadline: deadline.getTime()},
//   (err, response) => {
//     console.log('Got error, code = ' + err.code +
//                 ', message = ' + err.message);
//   });
