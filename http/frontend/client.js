require('google-protobuf')

const {BaseRequest, BaseResponse,
    GetLinksRequest, GetLinksResponse,
    PutLinkRequest, PutLinkResponse,
    RemoveLinkRequest, RemoveLinkResponse} = require('./drop_pb.js');
const {DropApiClient} = require('./drop_grpc_web_pb.js');

var client = new DropApiClient('http://' + window.location.hostname + ':8081', null, null);

console.log("hi " + process.env.NODE_ENV)
console.log("wow something")

// simple unary call
var request = new GetLinksRequest();
request.setBase(new BaseRequest());

console.log("about to call getLinks");
var linksStream = client.getLinks(request, {});

addElement = function(text, href) {
    let linksList = document.getElementById("linksList");
    const anchorTag = document.createElement("a");
    anchorTag.text = text;
    anchorTag.href = href;

    const deleteTag = document.createElement("a");
    deleteTag.text = "[X]";
    deleteTag.href = "javascript:";
    deleteTag.onclick = function(e) {
        var request = new RemoveLinkRequest();
        request.setBase(new BaseRequest());
        request.setLinktext(anchorTag.text);

        client.removeLink(request, {}, (err, response) => {
            }
        );
        document.getElementById("linksList").removeChild(e.target.parentNode);
    };

    const space = document.createElement("span")
    space.text = " ";

    const listElement = document.createElement("li");
    listElement.appendChild(deleteTag);
    listElement.appendChild(space);
    listElement.appendChild(anchorTag);

    linksList.prepend(listElement);
}

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
        request.setBase(new BaseRequest());
        var href = document.getElementById("linkHref").value;
        var text = document.getElementById("linkText").value;
        request.setLinkaddress(document.getElementById("linkHref").value);
        request.setLinktext(document.getElementById("linkText").value);

        addElement(text, href);

        client.putLink(request, {}, (err, response) => {
            }
        );
    }
}


linksStream.on('data', (response) => {
    addElement(response.getLinktext(), response.getLinkaddress());
});
    
console.log("after call");

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
