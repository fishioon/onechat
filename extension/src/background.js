const {chatNewClient, chatConn, chatPubmsg, chatJoinGroup} = require('./chat.js')

const token = 'ifishjin-123456';
const meta = {'Authorization': 'Bearer ' + token};

var client = chatNewClient('https://onechat.fishioon.com:1443')

chatConn(client, meta, token, (msg, err) => {
		if (err) {
			console.error(err)
		} else {
			console.log('recvmsg', msg)
			showmsg(lastTabid, msg)
		}
	}
)

var myTabs = []
var lastTabid = 0

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
