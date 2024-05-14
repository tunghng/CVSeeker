
import { useState } from "react"

import { FileUploader } from "react-drag-drop-files";
import LinkedinUploadInput from "../components/LinkedinUploadInput/LinkedinUploadInput";
import FeatherIcon from 'feather-icons-react'
import fileicon from '../assets/images/file.png'
import pdfToText from 'react-pdftotext'

const fileTypes = ["PDF"];

const UploadPage = () => {
    // ====== State Management ======
    const [urlInput, setUrlInput] = useState('')
    const [files, setFiles] = useState([])
    const [isDragging, setIsDragging] = useState(false)
    const [error, setError] = useState('')

    // ====== Event Handlers ======
    const linkedinUploadKeyDownHandler = (e) => {
        if (e.key === 'Enter' && urlInput.trim() !== '') {
            console.log(urlInput.trim())
        }
    }
    const linkedinUploadClickHandler = () => {
        if (urlInput.trim() !== '') {
            console.log(urlInput.trim())
        }
    }

    const handleChange = async (fileList) => {
        setError('')

        const filesArray = Array.isArray(fileList) ? fileList : Object.values(fileList);

        const invalidFiles = filesArray.filter(file => !fileTypes.includes(file.name.split('.').pop().toUpperCase()));

        if (invalidFiles.length > 0) {
            setError('Invalid file type. Please upload only PDF files.');
            return;
        }

        const updatedFiles = [...files, ...filesArray];
        setFiles(updatedFiles);
        await readPdfFiles(filesArray[0]);
    };

    const handleDragStateChange = (dragging) => {
        setIsDragging(dragging);
    };

    const removeFile = (index) => {
        const updatedFiles = [...files];
        updatedFiles.splice(index, 1);
        setFiles(updatedFiles);
    };

    const readPdfFiles = async (file) => {
        pdfToText(file)
            .then(text => console.log(text))
            .catch(error => console.error("Failed to extract text from pdf"))
    };

    return (
        <main>
            <div className="my-container-medium py-6">
                <h1 className="text-2xl font-bold">Upload profile</h1>

                {/* ====== Upload by link profile ====== */}
                <h2 className="mt-4 text-lg text-text">Upload profile by Linkedin Url</h2>
                <LinkedinUploadInput
                    value={urlInput}
                    onChange={(e) => setUrlInput(e.target.value)}
                    onPressEnter={linkedinUploadKeyDownHandler}
                    onClickButton={linkedinUploadClickHandler}
                />


                {/* ====== Upload PDF file ====== */}
                <h2 className="mt-8 text-lg text-text">Upload PDF file</h2>
                <FileUploader
                    handleChange={handleChange}
                    name="file"
                    multiple={true}
                    types={fileTypes}
                    onTypeError={() => setError('Invalid file type. Please upload only PDF files.')}
                    children={<CustomFileUploader isDragging={isDragging} />}
                    dropMessageStyle={{ display: "none" }}
                    onDraggingStateChange={handleDragStateChange}
                />
                <div className="mt-2 flex justify-between text-subtitle">
                    <p>Support formats: PDF</p>
                    <p>Maximum size: 1MB</p>
                </div>

                {/* ====== Error Message ====== */}
                {error && <p className="text-red-500 mt-2">{error}</p>}

                {/* ====== Uploaded files ====== */}
                <div className="flex flex-col">
                    {
                        files.length > 0 && (
                            <>
                                {files.map((file, index) => (
                                    <div key={index} className="mt-6 px-4 py-3 flex items-center rounded-xl bg-disable-light">
                                        <FeatherIcon icon="file" className="w-8 h-8 text-text " strokeWidth={1.8} />
                                        <div className="ml-2">
                                            <h3 className="text-lg text-text">{file.name}</h3>
                                            <p className="text-subtitle">{file.size} bytes</p>
                                        </div>
                                        <FeatherIcon icon="x" className="ml-auto w-6 h-6 text-text cursor-pointer" strokeWidth={1.8} onClick={() => removeFile(index)} />
                                    </div>
                                ))}
                                <button className="my-button my-button-primary self-end flex px-3 py-2 mt-4">
                                    <FeatherIcon icon="upload" strokeWidth={1.8} />
                                    <span className="ml-2 font-semibold">Upload</span>
                                </button>
                            </>
                        )
                    }
                </div>
            </div>
        </main>
    )
}

function CustomFileUploader({ isDragging }) {
    return (
        <div className={`mt-2 py-8 flex flex-col items-center border-2 ${isDragging ? 'border-primary bg-[#f2f7ff]' : 'border-border'} border-dashed rounded-xl transition-all duration-300 ease-in-out`}>
            <div className="relative">
                <img src={fileicon} alt="file icon" className="w-20 " />
                <FeatherIcon icon="upload" className={`absolute -right-3 -bottom-2 w-8 h-8 p-1.5 rounded-full ${isDragging ? 'bg-primary' : 'bg-title'} text-white transition-all duration-300 ease-in-out`} strokeWidth={1.9} />
            </div>

            <h1 className="mt-6">Drag and drop files here or
                <span className="ml-1 underline cursor-pointer underline-offset-2">Choose files</span>
            </h1>
        </div>
    )
}

export default UploadPage
