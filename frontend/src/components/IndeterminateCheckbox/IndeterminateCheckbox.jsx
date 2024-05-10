
import { useEffect, useRef } from 'react';
import FeatherIcon from 'feather-icons-react'

const IndeterminateCheckbox = ({ searchResults, setSearchResults }) => {
    
    const checkboxRef = useRef(null);
    const isAllSelected = searchResults.every(item => item.selected)
    const isIndeterminate = searchResults.some(item => item.selected) && !isAllSelected

    const handleSelectAll = () => {
        const updatedItems = searchResults.map(item => ({ ...item, selected: !isAllSelected }))
        setSearchResults(updatedItems)
    }

    useEffect(() => {
        if (checkboxRef.current) {
            checkboxRef.current.indeterminate = isIndeterminate;
        }        
    }, [isIndeterminate]);

    return (
        <div className='flex relative'>
            <input
                type="checkbox"
                ref={checkboxRef}
                checked={isAllSelected}
                onChange={handleSelectAll}
                className='cursor-pointer w-6 h-6 peer appearance-none
                        bg-white border-2 border-border rounded-[4px]
                        indeterminate:bg-primary indeterminate:border-primary
                        checked:bg-primary checked:border-primary
                        '
            />
            <FeatherIcon icon="check" className={`${isAllSelected ? 'block' : 'hidden'} absolute top-1/2 -translate-x-1/2 left-1/2 -translate-y-1/2 w-5 h-5 text-white pointer-events-none`} strokeWidth={3}/>
            <FeatherIcon icon="minus" className="absolute top-1/2 -translate-x-1/2 left-1/2 -translate-y-1/2 w-5 h-5 text-white hidden peer-indeterminate:block pointer-events-none" strokeWidth={3}/>
        </div>
    );

}

export default IndeterminateCheckbox
