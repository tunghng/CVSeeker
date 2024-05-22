
const DeleteThreadModal = ({ onClose, onDelete }) => {

    return (
        <div className="fixed top-0 left-0 w-screen h-screen z-50 flex items-center justify-center"
            onClick={onClose}
        >
            <div className="bg-white w-full max-w-80 mb-32 px-6 pt-4 pb-5 rounded-xl border border-border drop-shadow-lg"
                onClick={(e) => e.stopPropagation()}
            >
                <h2 className="text-xl font-semibold mb-3">Delete Thread</h2>
                <p className="text-sm text-subtitle mb-4">Are you sure you want to delete this thread?</p>

                <div className="flex justify-end space-x-4">
                    <button onClick={onClose} className="w-full my-button my-button-secondary">
                        Cancel
                    </button>
                    <button onClick={onDelete} className="w-full my-button my-button-danger py-1.5">
                        Delete
                    </button>
                </div>
            </div>
        </div>
    )
}

export default DeleteThreadModal
