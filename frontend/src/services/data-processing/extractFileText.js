
import pdfToText from 'react-pdftotext'

export default async function extractPdfFile(file) {
    
    try {
        const text = await pdfToText(file)
        return text
    }
    catch (error) {
        console.error("Failed to extract text from pdf", file.name, error)
    }
}

