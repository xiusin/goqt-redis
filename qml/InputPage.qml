import QtQuick 2.0
import QtQuick.Controls 2.12

Page {
    id: page
    // 定义参数，每行的高度
    property int rowHeight: 40
    // 定义参数，每行中，每列的间距
    property int rowSpacing: 8
    // 定义一列
    Column{
        id: column
        // 充满父类Page类
        anchors.fill: parent
        // 定义Column中，每行Row的间距
        spacing: 10
        Row{
            width: parent.width
            height: rowHeight
            spacing: rowSpacing
            // Row水平居中显示
            anchors.horizontalCenter: parent.horizontalCenter
            Label{
                id: lab
                text: "匹配表达式"
                // 定义垂直居中显示
                verticalAlignment: parent.verticalAlignment
                // 显示字符，水平靠右显示
                horizontalAlignment: Text.AlignRight
                // 设置宽度，Row的宽度的0.3
                width: 60
                height: parent.height
            }

            TextField{
                id: patternInput
                placeholderText: "回车查看所有影响Key"
                width: parent.width - 100
                height: parent.height
                onEditingFinished: {
                    ctxObject.onEditingFinished(patternInput.text)
                    affectedKeys.text = ctxObject.patternKey
                }
            }
        }

        Row{
              width: parent.width
              height: 280
              anchors.horizontalCenter: parent.horizontalCenter
              ScrollView {
                  anchors.fill: parent
                  TextArea{
                        id: affectedKeys
                        width: parent.width
                        height: parent.height
                        text: ""
                  }
              }
        }

        Row{
            width: parent.width
            height: rowHeight
            spacing: rowSpacing
            anchors.horizontalCenter: parent.horizontalCenter
            Button{
                id: b2
                text: "提交删除"
                width: parent.width * 0.15
                height: parent.height
                onClicked: {
                    batchDeleteWin.visible = false
                    ctxObject.onClicked(patternInput.text)
                }
            }
        }
    }
}
