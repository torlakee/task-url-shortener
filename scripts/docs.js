window.hasClass = function (el, className) {
    if (el.classList) {
        return el.classList.contains(className);
    } else {
        return new RegExp('(^| )' + className + '( |$)', 'gi').test(el.className);
    }
}

window.removeClass = function (el, className) {

    if (el.classList)
        el.classList.remove(className);
    else
        el.className = el.className.replace(new RegExp('(^|\\b)' + className.split(' ').join('|') + '(\\b|$)', 'gi'), ' ');
}

window.addClass = function (el, className) {

    if (el.classList)
        el.classList.add(className);
    else
        el.className += ' ' + className;
}

let methodTabs = document.querySelectorAll('.tab');
let methodTestButtons = document.querySelectorAll('.post-test-button');

methodTabs.forEach((tab) => {
    tab.addEventListener('click', () => {

        if (!hasClass(tab, 'on')) {

            let holder = tab.closest('.method');

            let activeTab = holder.querySelector('.tab.on');
            let activeSection = holder.querySelector('.tab-section.on');

            if (activeTab != null) {
                if (hasClass(activeTab, 'on')) {
                    removeClass(activeTab, 'on');
                    removeClass(activeSection, 'on');
                }
            }
            let page = holder.querySelector('.tab-section[data-tab="' + tab.dataset.tab + '"]');

            addClass(tab, 'on');
            addClass(page, 'on');
        }
    });
});

methodTestButtons.forEach((button) => {
    button.addEventListener('click', () => {

        let holder = button.closest('.tab-section.on');

        let data = holder.querySelector('.param').value;

        if (data.length !== 0) {
            fetch('/', {
                method: 'post',
                body: data,
                mode: 'cors',
                headers: new Headers({
                    'content-type': 'text/plain',
                    'charset': 'utf-8'
                })
            })
                .then(response => {
                    let responseType = response.headers.get("content-type");

                    let resultHolder = holder.querySelector('.result');

                    if (responseType && responseType.indexOf("application/json") !== -1) {
                        response.text().then(v => {
                            resultHolder.innerHTML = `<div style="text-align: center; margin:0 auto;width:200px;"><div class="result-caption">Резултат</div></div><div class="error"><div class="caption">Възникна грешка</div><div class="content">${v}</div></div>`;
                        });
                    } else {


                        response.text().then(v => {
                            resultHolder.innerHTML = `<div style="text-align: center; margin:0 auto;width:200px;"><div class="result-caption">Резултат</div></div><div class="result-output"><div class="caption">Кратък адрес</div><div class="content">${v}</div></div>`;
                        });
                    }
                });

        } else {


        }
    });
});