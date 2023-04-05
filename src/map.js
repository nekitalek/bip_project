    // Функция ymaps.ready() будет вызвана, когда
    // загрузятся все компоненты API, а также когда будет готово DOM-дерево.
    ymaps.ready(init);
    function init(){
        // Создание карты.
        var myMap = new ymaps.Map("map", {
            center: [59.939098, 30.315868],
            zoom: 11,
            type: 'yandex#map',
            controls: ['zoomControl', 'geolocationControl']
        }, {
            searchControlProvider: 'yandex#search'
        });
    }

