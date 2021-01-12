import QtQuick 2.12
import QtQuick.Window 2.12

Window {
    id: batchDeleteWin
    visible: true
    width: 600
    height: 400
    title: qsTr("批量删除Key")
    InputPage{
        // 充满父类
        anchors.fill: parent
        // 设置margins
        anchors.margins: 10
    }
}
