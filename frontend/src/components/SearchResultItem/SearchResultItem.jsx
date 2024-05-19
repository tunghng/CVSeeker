
import FeatherIcon from 'feather-icons-react'

const SearchResultItem = ({ item, viewMode, onSelectClick, onDetailClick, onDownloadClick }) => {

    const handleSelectClick = () => {
        onSelectClick(item.id);
    }

    const handleDetailClick = (e) => {
        e.stopPropagation();
        onDetailClick(item);
    }

    const handleDownloadClick = (e) => {
        e.stopPropagation();
        onDownloadClick(item);
    }

    return (
        <div className="h-14 flex border-b border-border hover:bg-blue-100/50"
            onClick={handleSelectClick}
        >
            <div className="w-12 flex justify-center items-center relative">
                <input
                    type="checkbox"
                    checked={item.selected}
                    readOnly
                    className="appearance-none peer w-6 h-6 border-2 border-border rounded-full
                             checked:bg-primary checked:border-primary"
                />
                <FeatherIcon icon="check" className='hidden peer-checked:block absolute top-1/2 -translate-x-1/2 left-1/2 -translate-y-1/2 w-[18px] h-[18px] text-white pointer-events-none' strokeWidth={3} />
            </div>

            <p className='w-44 my-auto text-left text-base font-semibold cursor-default truncate'>{item.basic_info.full_name}</p>

            <p className='w-32 my-auto text-center text-sm cursor-default truncate'>{item.basic_info.point}</p>

            <p className='flex-1 my-auto text-center text-sm cursor-default truncate'>{item.basic_info.education_level}</p>

            <div className="w-20 sm:w-48 flex items-center justify-center gap-x-1">
                <button
                    className='px-2.5 py-2 sm:py-1 bg-disable-light text-sm rounded-full flex items-center gap-x-1 hover:bg-secondary-subtle'
                    onClick={handleDetailClick}>
                    <FeatherIcon icon="info" className='w-4 h-4' strokeWidth={1.6} />
                    <p className='hidden sm:block'>Detail</p>
                </button>
                <button
                    className='px-2.5 py-2 sm:py-1 bg-disable-light text-sm rounded-full flex items-center gap-x-1 hover:bg-secondary-subtle'
                    onClick={handleDownloadClick}>
                    <FeatherIcon icon="download" className='w-4 h-4' strokeWidth={1.6} />
                    <p className='hidden sm:block'>Download</p>
                </button>
            </div>
        </div>
    );
}

export default SearchResultItem
