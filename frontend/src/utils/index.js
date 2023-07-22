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
