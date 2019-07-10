const {ChatClient} = require('../chat/chat_grpc_web_pb')
const {ConnReq, GroupActionReq, Msg, PubMsgReq} = require('../chat/chat_pb')

var myTabs = [];
var lastTabid = 0;

var client = new ChatClient('https://onechat.fishioon.com:1443', null, null)

var connReq = new ConnReq();
connReq.setToken('ifishjin-123456');

var stream = client.conn(connReq, {})
stream.on('data', (rsp) => {
	console.log(rsp.getFromId(), rsp.getToId(), rsp.getContent())
})

// Toolbar Button
chrome.browserAction.onClicked.addListener(function () {
	pubmsg(myTabs[lastTabid].groupid, 'hello')
})

chrome.tabs.onUpdated.addListener(function (tabId, _, tab) {
	lastTabid = tabId;
	if (myTabs[tabId] === undefined || myTabs[tabId].url !== tab.url) {
		myTabs[tabId] = {url: tab.url, groupid: 0}
		joinGroup(tab.url, tabId)
	}
})

chrome.commands.onCommand.addListener(function (command) {
	console.log('Command:', command)
	showmsg('command:' + cmd)
})

function joinGroup(url) {
	var req = new GroupActionReq()
	req.setAction('join')
	req.setGid(url)
	client.groupAction(req, {}, (err, rsp) => {
		if (err) {
			console.error('json group', url, err)
		} else {
			console.log(rsp)
		}
	})
}

function pubmsg(gid, content) {
	var msg = new Msg()
	msg.setToId(gid)
	msg.setContent(content)
	var req = new PubMsgReq()
	req.setMsg(msg)
	client.pubMsg(req, {}, (err, rsp) => {
		if (err) {
			console.error('pub msg', msg, err)
		} else {
			console.log(rsp)
		}
	})
}

function showmsg(tabid, msg) {
	chrome.tabs.sendMessage(tabid, {
		type: 'danmu',
		message: {user: {name: msg}, content: {text: msg, image: ''}}
	});
}
