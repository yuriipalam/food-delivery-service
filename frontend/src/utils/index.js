// truncates text
export function truncate(str, limit) {
    if (str.length > limit) {
        return str.substring(0, limit) + '...';
    }
    return str;
}

// gets full height of any element including all the space-adders
export function getElmHeight(node) {
    const list = ['margin-top', 'margin-bottom', 'border-top', 'border-bottom', 'padding-top', 'padding-bottom', 'height']

    const style = window.getComputedStyle(node)
    return list
        .map(k => parseInt(style.getPropertyValue(k), 10))
        .reduce((prev, cur) => prev + cur)
}

// just formats date to EU format
export function formatDate(dateStr) {
    return new Date(dateStr).toLocaleString('en-GB');
}

// scroll to the .explore block when opening a page
export function scrollToExploreBlock() {
    const scrollTo = document.querySelector('.explore').offsetTop - 40

    window.scrollTo({
        top: scrollTo,
        behavior: 'smooth'
    })
}

// setting height of main block
// it's needed when we use search and less/no results found
// that our page won't change its height
export function setMainHeight() {
    const main = document.querySelector('main')
    main.style.minHeight = main.offsetHeight + 'px'
}
