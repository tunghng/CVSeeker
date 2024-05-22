
export default function generateThreadName() {
    const now = new Date();
    const vietnamTimeOffset = 7 * 60;
    const localTimeOffset = now.getTimezoneOffset();
    const vietnamTime = new Date(now.getTime() + (vietnamTimeOffset + localTimeOffset) * 60000);

    const year = vietnamTime.getFullYear();
    const month = String(vietnamTime.getMonth() + 1).padStart(2, '0');
    const day = String(vietnamTime.getDate()).padStart(2, '0');
    const hours = String(vietnamTime.getHours()).padStart(2, '0');
    const minutes = String(vietnamTime.getMinutes()).padStart(2, '0');
    const seconds = String(vietnamTime.getSeconds()).padStart(2, '0');

    const timeStr = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return timeStr;
}