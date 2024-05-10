
import { useRef, useState } from "react"

import { FileUploader } from "react-drag-drop-files";
import FeatherIcon from 'feather-icons-react'
import fileicon from '../assets/images/file.png'

const fileTypes = ["PDF"];

const UploadPage = () => {
    // ====== State Management ======
    const [urlInput, setUrlInput] = useState('')
    const urlInputDOM = useRef(null)

    const [file, setFile] = useState(null);
    const [isDragging, setIsDragging] = useState(false);

    // ====== Event Handlers ======
    const handleChange = (file) => {
        setFile(file);
    };

    const handleDragStateChange = (dragging) => {
        setIsDragging(dragging);
    };

    return (
        <main>
            <div className="my-container-medium py-6">
                <h1 className="text-2xl font-bold">Upload profile</h1>

                {/* ====== Upload by link profile ====== */}
                <h2 className="mt-4 text-lg text-text">Upload profile by Linkedin Url</h2>
                <input
                    type="text"
                    className="w-full mt-2 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-text font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out"
                    placeholder="Paste Linkedin URL here..."
                    value={urlInput}
                    onChange={(e) => setUrlInput(e.target.value)}
                    // onKeyDown={searchKeyDownHandler}
                    ref={urlInputDOM}
                />


                {/* ====== Upload PDF file ====== */}
                <h2 className="mt-6 text-lg text-text">Upload PDF file</h2>
                <FileUploader
                    handleChange={handleChange}
                    name="file"
                    types={fileTypes}
                    hoverTitle="Drop file"
                    children={<CustomFileUploader isDragging={isDragging} />}
                    dropMessageStyle={{ display: "none" }}
                    onDraggingStateChange={handleDragStateChange}
                />
                <div className="mt-2 flex justify-between text-subtitle">
                    <p>Support formats: PDF</p>
                    <p>Maximum size: 1MB</p>
                </div>

                {/* ====== Uploaded files ====== */}
                <div>
                    {file && (
                        <div className="mt-6 px-4 py-3 flex items-center rounded-xl bg-disable-light">
                            <FeatherIcon icon="file" className="w-8 h-8 text-text " strokeWidth={1.8} />
                            <div className="ml-2">
                                <h3 className="text-lg text-text">{file.name}</h3>
                                <p className="text-subtitle">{file.size} bytes</p>
                            </div>
                            <FeatherIcon icon="x" className="ml-auto w-6 h-6 text-text cursor-pointer" strokeWidth={1.8} onClick={() => setFile(null)} />
                        </div>
                    )}
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

            <h1 className="mt-6">Drag and drop file here or
                <span className="ml-1 underline cursor-pointer underline-offset-2">Choose file</span>
            </h1>
        </div>
    )
}

export default UploadPage
