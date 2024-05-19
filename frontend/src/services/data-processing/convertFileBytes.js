
export default async function convertFileBytes(file) {

    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        
        reader.onload = () => {
            const arrayBuffer = reader.result;
            const bytes = new Uint8Array(arrayBuffer);
            
            const decoder = new TextDecoder("utf-8");
            const str = decoder.decode(bytes);
            
            resolve(str);
        };
        
        reader.onerror = () => {
            reject(new Error("Failed to read file"));
        };
        
        reader.readAsArrayBuffer(file);
    });
}
