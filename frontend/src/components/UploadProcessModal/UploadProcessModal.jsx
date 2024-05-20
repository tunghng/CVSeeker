
import FeatherIcon from "feather-icons-react"

const UploadProcessModal = ({ showUploadProcessModal, uploadedFiles, onClose, onDetail }) => {
    return (
        <div className={`${showUploadProcessModal ? 'block' : 'hidden'} fixed top-0 left-0 w-screen min-h-screen z-20 bg-black/40 flex items-center justify-center`}
            onClick={onClose}>
            <div className="bg-background my-container-small min-h-96 max-h-[30rem] overflow-y-scroll rounded-xl py-6 no-scrollbar"
                onClick={(e) => e.stopPropagation()}>
                <h2 className="text-lg font-semibold">Uploaded Files</h2>

                <div>
                    {uploadedFiles.length > 0 ? (
                        <div className="mt-4">
                            {uploadedFiles.map((file, index) => (
                                <div key={index} className={`mt-4 px-3 py-2 flex justify-between items-center rounded-xl 
                                    ${file.status === "Processing" && `bg-disable-light`}
                                    ${file.status === "Failed" && `bg-red-100`}
                                    ${file.status === "Success" && `bg-green-200/50`}`}>
                                    <div>
                                        <h3 className={`text-lg font-semibold
                                            ${file.status === "Processing" && `text-text`}
                                            ${file.status === "Failed" && `text-red-500`}
                                            ${file.status === "Success" && `text-green-600/90`}
                                        `}>
                                            {file.name === "" ? "Pending" : file.name}
                                        </h3>
                                        <p className="text-subtitle">
                                            Status: {file.status}
                                        </p>
                                    </div>

                                    {file.status === "Success" && (
                                        <button className="my-button bg-green-600/80 text-white flex items-center py-1 rounded-full"
                                            onClick={() => onDetail(file)}>
                                            <FeatherIcon icon="info" className="w-5 h-5 mr-1" strokeWidth={2} />
                                            Detail
                                        </button>
                                    )}
                                </div>
                            ))}
                        </div>
                    )
                        :
                        <p className="mt-4 text-text text-center">No files uploaded</p>
                    }
                </div>
            </div>
        </div>
    )
}

export default UploadProcessModal