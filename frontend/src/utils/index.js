import {computed} from "vue";

export function truncate(str, limit) {
    if (str.length > limit) {
        return str.substring(0, limit) + '...';
    }
    return str;
}
