import QtQuick 2.2
import QtQuick.Window 2.13
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.1

Window {
    onClosing: {
        ctxObject.onClosing()
    }
    id: addServerWindow
    visible: true
    flags: Qt.Dialog
    modality: Qt.WindowModal
    width: 400
    height: 250
    minimumHeight: 250
    minimumWidth: 400
    title: "添加新服务器"

    GridLayout {
        columns: 2
        anchors.fill: parent
        anchors.margins: 10
        rowSpacing: 10
        columnSpacing: 10

        Label {
            text: "名称    "
        }
        TextField {
            id: name
            text: ""
            Layout.fillWidth: true
        }

        Label {
            text: "地址    "
        }
        TextField {
            id: address
            text: ""
            Layout.fillWidth: true
        }

        Label {
            width: 150
            text: "端口    "
        }
        TextField {
            id: port
            text: ""
            Layout.fillWidth: true
        }


        Label {
            width: 150
            text: "密码    "
        }
        TextField {
            id: password
            text: ""
            echoMode: TextInput.Password
            Layout.fillWidth: true
        }

        Item {
            Layout.columnSpan: 2
            Layout.fillWidth: true
            implicitHeight: button.height
            GridLayout {
                anchors.centerIn: parent
                columns: 2
                Button {
                    onClicked: {
                        ctxObject.testServer(name.text, address.text, port.text, password.text)
                    }
                    contentItem: Text {
                        text: "测试"
                        color: "white"
                        font.pixelSize: 14
                        font.family: "Arial"
                        font.weight: Font.Thin
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }

                    background: Rectangle {
                        implicitWidth: 60
                        implicitHeight: 30
                        color: "#0ACF97"
                        radius: 0
                        layer.enabled: true
                        layer.effect: DropShadow {
                            transparentBorder: true
                            color: "#ffffff"
                            samples: 20
                        }
                    }
                }

                Button {
                    contentItem: Text {
                        text: "保存"
                        color: "white"
                        font.pixelSize: 14
                        font.family: "Arial"
                        font.weight: Font.Thin
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }

                    background: Rectangle {
                        implicitWidth: 60
                        implicitHeight: 30
                        color: "#212730"
                        radius: 0
                        layer.enabled: true
                        layer.effect: DropShadow {
                            transparentBorder: true
                            color: "#ffffff"
                            samples: 20
                        }
                    }
                    onClicked: {
                        ctxObject.saveServer(name.text, address.text, port.text, password.text)
                        addServerWindow.close();
                    }
                }
            }
        }
    }
}
