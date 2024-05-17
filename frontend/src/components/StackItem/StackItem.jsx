
import FeatherIcon from "feather-icons-react";

const StackItem = ({ item, onDetailClick, onRemoveClick, showRemoveIcon }) => {

    const handleDetailClick = (e) => {
        e.stopPropagation();
        onDetailClick(item);
    }

    const handleRemoveClick = (e) => {
        e.stopPropagation();
        onRemoveClick(item.id);
    }

    return (
        <div
            className="px-3 py-2 mt-2 flex items-center justify-between rounded-md bg-disable-light cursor-pointer"
            onClick={handleDetailClick}
        >
            <p>{item.basic_info.full_name}</p>

            {showRemoveIcon && (
                <FeatherIcon
                    icon="x"
                    className="w-6 h-6 p-1 rounded-full cursor-pointer hover:bg-subtitle"
                    onClick={handleRemoveClick}
                />
            )}
        </div>
    )
}

export default StackItem
