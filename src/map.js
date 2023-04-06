    // Функция ymaps.ready() будет вызвана, когда
    // загрузятся все компоненты API, а также когда будет готово DOM-дерево.
    function init() {
        var myMap = new ymaps.Map('map', {
            center: [59.939098, 30.315868],
            zoom: 11,
            controls: []
        });
        
        // Создадим экземпляр элемента управления «поиск по карте»
        // с установленной опцией провайдера данных для поиска по организациям.
        var searchControl = new ymaps.control.SearchControl({
            options: {
                provider: 'yandex#search'
            }
        });
        
        // myMap.controls.add(searchControl);
        
        // Программно выполним поиск определённых кафе в текущей
        // прямоугольной области карты.
        //searchControl.search('Шоколадница');
    } 
    ymaps.ready(init);

      
