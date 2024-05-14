
import pdfToText from 'react-pdftotext'

const extractPdfFile = async (file) => {
    try {
        const text = await pdfToText(file)
        return text
    }
    catch (error) {
        console.error("Failed to extract text from pdf", file.name, error)
    }
}

export default extractPdfFile