let clockID = 0;

$().ready(function () {

    // 验证表单
    $('#mainForm')
        .form({
            inline: 'true',
            on: 'blur',
            fields: {
                kafkaIP: {
                    rules: [
                        {
                            type: 'empty',
                            prompt: 'kafka IP:Port不能为空'
                        },
                        {
                            type: 'minLength[12]',
                            prompt: '请输入正确的IP:端口,多个IP:端口 使用英文逗号分隔'
                        },
                    ]
                },
                messageContent: {
                    rules: [
                        {
                            type: 'empty',
                            prompt: '消息内容不能为空'
                        }
                    ]
                },
                messageNumber: {
                    rules: [
                        {
                            type: 'empty',
                            prompt: '发送次数不能为空'
                        },
                        {
                            type: 'integer[1..60]',
                            prompt: '入发送次数需要在1 - 60 之间'
                        }
                    ]
                },
                kafkaTopic: {
                    rules: [
                        {
                            type: 'empty',
                            prompt: 'Topic不能为空'
                        }
                    ]
                }
            },
            onSuccess: function () {
                mainStart();
            },
            onFailure: function () {
                console.log("数据验证失败")
            }

        });

    // 重置按钮
    $("#getReset").click(function () {
        location.reload(true);
    });

    // 阻止表单默认提交
    $('#mainForm').on('submit', function () {
        event.preventDefault()
    })
});

// 开始检测
function mainStart() {

    $("#getStart").attr("class", "ui disabled button");
    $("#getStart").html("检测中...");

    var times = $("#messageNumber").val();

    $("#mainForm").ajaxSubmit({
        success: function (data) {
            clockID = setInterval(streamData, 500, data);
            // setTimeout(window.clearInterval, 15000, clockID);
        },
        error: function (message) {
            alert("联系后端服务失败");
        }

    });

}


function streamData(data) {
    $.ajax({
        type: 'get',
        url: '/getStreamData',
        success: function (res) {

            if (res != "Empty") {
                $("#output1").append(res + "<br>");
            }
        },
        error: function (err) {
            $("#output1").append(data + "<br>");
            console.log(err);
        }
    });

}

