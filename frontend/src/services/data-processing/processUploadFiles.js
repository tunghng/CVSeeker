
import extractPdfFile from "./extractFileText";
import convertFileBytes from "./convertFileBytes";
import { v4 as uuidv4 } from 'uuid';

export default async function processUploadFiles(files) {

    const processedFiles = await Promise.all(files.map(async (file) => {
        const content = await extractPdfFile(file);
        const fileBytes = await convertFileBytes(file);
        const uuid = uuidv4();
        const name = file.name;
        return { content, fileBytes, uuid, name };
    }));

    return processedFiles;
}
