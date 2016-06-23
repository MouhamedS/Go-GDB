import QtQuick 2.0
import QtQuick.Controls 1.1

Rectangle {
    id: root
    color: "#ffffff"

    width: 320
    height: 320

    TextEdit {
        id: textEdit1
        x: 8
        y: 8
        width: 304
        height: 20
        text: ctrl.name
        font.pixelSize: 12
    }

    Button {
        id: button1
        x: 8
        y: 34
        width: 304
        height: 27
        text: qsTr("Button")
        onClicked: {
            ctrl.hello()
        }
    }

    Text {
        id: text1
        x: 8
        y: 67
        width: 304
        height: 23
        text: ctrl.message
        font.pixelSize: 12
    }
Binding { target: ctrl; property: "name"; value: textEdit1.text }
}


