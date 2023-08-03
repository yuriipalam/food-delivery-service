// truncates text
export function truncate(str, limit) {
    if (str.length > limit) {
        return str.substring(0, limit) + '...';
    }
    return str;
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

// hide footer
export function hideFooter() {
    const footer = document.querySelector('footer')
    footer.style.display = 'none'
}

// show footer
export function showFooter() {
    const footer = document.querySelector('footer')
    footer.style.display = 'block'
}
