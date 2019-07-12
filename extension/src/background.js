const {ChatClient} = require('./chat_grpc_web_pb')
const {ConnReq, GroupActionReq, Msg, PubMsgReq} = require('./chat_pb')

var myTabs = []
var lastTabid = 0

const token = 'ifishjin-123456';
const meta = {'Authorization': 'Bearer ' + token};

var client = new ChatClient('https://onechat.fishioon.com:1443', null, null)
var connReq = new ConnReq()
connReq.setToken(token)
var stream = client.conn(connReq, {})
stream.on('data', (rsp) => {
	console.log('conn data', rsp)
	var msg = {
		fromId: rsp.getFromId(),
		toId: rsp.getToId(),
		content: rsp.getContent(),
	}
	showmsg(lastTabid, msg)
})
stream.on('status', status => {
	console.log('status', status)
})
stream.on('error', err => {
	console.log('error', err)
})
stream.on('end', () => {
	console.log('end', 'stream end signal received')
})

chrome.browserAction.onClicked.addListener(function () {
	chatPubmsg(client, meta, myTabs[lastTabid].groupid, 'hello')
})

chrome.tabs.onUpdated.addListener(function (tabId, _, tab) {
	lastTabid = tabId;
	if (myTabs[tabId] === undefined || myTabs[tabId].url !== tab.url) {
		myTabs[tabId] = {url: tab.url, groupid: 0}
		chatJoinGroup(client, meta, tab.url, tabId)
	}
})

chrome.commands.onCommand.addListener(function (command) {
	console.log('Command:', command)
})

function showmsg(tabid, msg) {
	chrome.tabs.sendMessage(tabid, {
		type: 'danmu',
		message: {user: {name: msg.fromId}, content: {text: msg.content, image: ''}}
	});
}

function chatJoinGroup(client, meta, url) {
	var req = new GroupActionReq()
	req.setAction('join')
	req.setGid(url)
	client.groupAction(req, meta, (err, rsp) => {
		if (err) {
			console.error('json group', url, err)
		} else {
			console.log(rsp)
		}
	})
}

function chatPubmsg(client, meta, gid, content) {
	var msg = new Msg()
	msg.setToId(gid)
	msg.setContent(content)
	var req = new PubMsgReq()
	req.setMsg(msg)
	client.pubMsg(req, meta, (err, rsp) => {
		if (err) {
			console.error('pub msg', msg, err)
		} else {
			console.log(rsp)
		}
	})
}
