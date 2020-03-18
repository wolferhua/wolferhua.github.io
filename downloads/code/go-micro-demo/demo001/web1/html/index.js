//alert("message")
(function ($) {

    const sessionKey = 'name';
    const host = window.location.host;
    const wsUrl = "ws://" + host + "/web1/ws?name=";

    let $messageBox = $(".message-box"); //消息框
    let $message = $("#message"); //消息内容
    let $sendBut = $("#send-but"); //消息按钮
    let $userList = $(".user-list");//用户列表

    let ws = null;


    // 获取用户名称
    let name = sessionStorage.getItem(sessionKey);// 如果存在

    // 加载用户列表
    let loadUserList = _.debounce(function () {
        //console.log($)
        $.ajax({
            url:'/web1/ws/users',
            dataType:'json',
            success:function (ret) {
                //console.log(ret);
                // 渲染消息数据
                let $warp = $('<div></div>');
                for (let i in ret){
                    $warp.append('<li class="list-group-item" data-name="'+ret[i]+'">'+ret[i]+'</li>')
                }
                $userList.html($warp.html());
            }
        })
    },3000);

    // 显示消息
    function writeMessage(content) {
        let item = $('<div class="col"></div>');
        item.html(content);
        writeMessageRow(item);
    }

    function writeMessageRow(content) {
        let item = $('<div class="row"></div>');
        item.html(content);
        $messageBox.append(item);
    }


    // 名称输入提示
    function namePrompt() {
        // 没有存储名称。
        do {
            name = window.prompt("请输入名称：");
            name = name.trim();
            // 存入本地存储
            sessionStorage.setItem(sessionKey, name);
        } while (!name)
    }

    // 链接服务器
    function connectServer() {
        ws = new WebSocket(wsUrl + name);

        // 事件监听
        ws.onopen = function (evt) {
            writeMessage('<p class="text-success text-center">与服务器链接已经建立。</p>');
            loadUserList();
        };
        ws.onclose = function (evt) {
            writeMessage('<p class="text-danger text-center">链接断开，请重新链接。</p>');
            ws = null;
        };
        ws.onmessage = function (evt) {
            let data = JSON.parse(evt.data);
            if (!data || !data.code || !data.message) return; // 消息体不匹配

            if (data.code >= 500) {
                // 系统异常消息：退出，并关闭链接
                sessionStorage.removeItem(sessionKey);
                writeMessage('<p class="text-danger text-center">' + data.message + '</p>');
                ws = null;
                return;
            }

            if (data.code >= 200) {
                // 用户消息。显示消息内容
                var message = '';
                if (data.name == name) {
                    //本人发送消息
                    message += '<div class="col-auto"><p class="text-success">' + data.name + '</p></div>'
                } else {
                    //其他人发送消息
                    message += '<div class="col-auto"><p class="text-muted">' + data.name + '</p></div>'
                }

                message += '<div class="col">' + data.message + '</div>';
                writeMessageRow(message);
                return;
            }


            // 系统通知消息
            if (data.code >= 100) {
                writeMessage('<p class="text-muted  text-center">' + data.message + '</p>');
                loadUserList();
                return;
            }

        };
        ws.onerror = function (evt) {
            writeMessage('<p class="text-danger text-center">服务器错误，请重新链接。</p>');
            ws = null;
        };
    }


    // 节流函数，防止重复点击。
    let send = _.debounce(function () {
        var content = $message.val().trim();
        if (!content) {
            alert('请输入内容！');
            return;
        }
        if (ws == null) {
            namePrompt();
            connectServer();
        }
        ws.send(content);
        $message.val('');
        $message.focus();
    }, 200);

    // 回车键发送消息
    $message.keypress(function (event) {
        if (event.keyCode === 13) {
            // 阻止事件冒泡
            if (event.preventDefault) event.preventDefault;
            event.returnValue = false;
            // 发送消息
            send();
            return false;
        }
    });
    // 按钮发送
    $sendBut.click(send);

    // 开始连接。
    if (name) {
        connectServer();
    } else {
        namePrompt();
        connectServer();
    }

})(jQuery);