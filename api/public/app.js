new Vue({
  el: '#app',

  data: {
    ws: null,
    newMsg: '',
    chatContent: '',
    email: null,
    username: null,
    joined: false
  },
  created: function() {
    var self = this;
    var wsUrl = 'ws://' + window.location.host + '/ws/' + getParam("id", window.location.href);
    this.ws = new WebSocket(wsUrl);
    this.ws.addEventListener('message', function(e) {
			var msg = JSON.parse(e.data);
			self.chatContent += '<div class="chip">'
				+ '<img src="' + self.gravatarURL(msg.email) + '">'
				+ msg.username
				+ '</div>'
				+ msg.message + '<br/>';

			var element = document.getElementById('chat-messages');
			element.scrollTop = element.scrollHeight;
    });
  },
  methods: {
    send: function () {
			if (this.newMsg != '') {
				this.ws.send(
					JSON.stringify({
						email: this.email,
						username: this.username,
						message: $('<p>').html(this.newMsg).text()
					}));
				this.newMsg = '';
			}
    },
    join: function () {
			if (!this.email) {
				Materialize.toast('You must enter an email', 2000);
				return
			}
			if (!this.username) {
				Materialize.toast('You must choose a username', 2000);
				return
			}
			this.email = $('<p>').html(this.email).text();
			this.username = $('<p>').html(this.username).text();
			this.joined = true;
    },
    gravatarURL: function(email) {
      return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
    }
  }
});

function getParam(name, url) {
	if (!url) url = window.location.href;
	name = name.replace(/[\[\]]/g, "\\$&");
	var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
			results = regex.exec(url);
	if (!results) return null;
	if (!results[2]) return '';
	return decodeURIComponent(results[2].replace(/\+/g, " "));
}
