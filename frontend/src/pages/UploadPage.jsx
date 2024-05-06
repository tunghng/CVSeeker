
import { useRef, useState } from "react"

const UploadPage = () => {
    // ====== State Management ======
    const [urlInput, setUrlInput] = useState('')
    const urlInputDOM = useRef(null)


    // ====== Event Handlers ======


    return (
        <main>
            <div className="my-container-medium py-6">
                <h1 className="text-2xl font-bold">Upload profile</h1>

                <input
                    type="text"
                    className="w-full mt-5 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-text font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out"
                    placeholder="Paste Linkedin URL here..." 
                    value={urlInput}
                    onChange={(e) => setUrlInput(e.target.value)}
                    // onKeyDown={searchKeyDownHandler}
                    ref={urlInputDOM}
                />

                <h2>Drag and drop your PDF here</h2>
                <p>or</p>
                <label htmlFor="file" className="my-button my-button-primary mt-5">Browse</label>
                <input type="file" name="file" id="file" className="hidden" />

            </div>
        </main>
    )
}

export default UploadPage