
const LoadingModal = ({ showLoadingModal }) => {

    return (
        <div className={`${showLoadingModal ? 'block' : 'hidden'} fixed top-0 left-0 w-screen min-h-screen z-50 bg-black/20 flex items-center justify-center`}>
            <div className="bg-white p-4 rounded-lg">
                <div className="flex justify-center items-center">
                    <div className="loader-2"></div>
                    <h1 className="ml-4 text-lg font-semibold">Starting chat session...</h1>
                </div>
            </div>
        </div>

    )
}

export default LoadingModal