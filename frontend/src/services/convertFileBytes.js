
export async function convertFileBytes(file) {
    const reader = new FileReader();
    return new Promise((resolve, reject) => {
        reader.onload = (event) => {
            resolve(event.target.result);
        };
        reader.onerror = (error) => {
            reject(error);
        };
        reader.readAsArrayBuffer(file);
    });
}