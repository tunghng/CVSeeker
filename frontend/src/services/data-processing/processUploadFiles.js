
import extractPdfFile from "./extractFileText";
import convertFileBytes from "./convertFileBytes";

export default async function processUploadFiles(files) {

    const processedFiles = await Promise.all(files.map(async (file) => {
        const content = await extractPdfFile(file);
        const fileBytes = await convertFileBytes(file);
        return { content, fileBytes };
    }));

    return processedFiles;
}
