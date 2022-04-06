var count = 1;

function ChangeColor(count) {

    var element = document.getElementById("mainText");
    var text = 'Текст изменен один раз';
    var text2 = 'Текст изменен ВТОРОЙ раз';
    var text3 = 'ТРЕТИЙ РАЗ';
    var text4 = 'И ПОССЛДЕНИЙ РАЗ ЧЕТВЕРТЫЙ';

    if (count == 1) {
        element.textContent = text;
        element.style.color = 'blue';
    }

    if (count == 2) {
        element.textContent = text2;
        element.style.color = 'red';
    }

    if (count == 3) {
        element.textContent = text3;
        element.style.color = 'yellow';
    }

    if (count == 4) {
        element.textContent = text4;
        element.style.color = 'black';
    }
}