
const RenameModal = ({ value, onChange, onClose, onRename }) => {

    return (
        <div className="fixed top-0 left-0 w-screen h-screen z-50 flex items-center justify-center"
            onClick={onClose}
        >
            <div className="bg-white w-full max-w-80 mb-32 px-6 pt-4 pb-5 rounded-xl border border-border drop-shadow-lg"
                onClick={(e) => e.stopPropagation()}
            >
                <h2 className="text-xl font-semibold mb-3">Rename Thread</h2>
                <input
                    type="text"
                    value={value}
                    onChange={onChange}
                    placeholder="Enter new name"
                    autoFocus
                    onKeyDown={(e) => e.key === 'Enter' && onRename()}
                    className="w-full p-2 border border-border rounded-md mb-4"
                />
                <div className="flex justify-end space-x-4">
                    <button onClick={onClose} className="w-full my-button my-button-secondary">
                        Cancel
                    </button>
                    <button onClick={onRename} className="w-full my-button my-button-primary py-1.5">
                        Rename
                    </button>
                </div>
            </div>
        </div>
    )
}
export default RenameModal