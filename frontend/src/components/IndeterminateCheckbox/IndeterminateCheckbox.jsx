
import { useEffect, useRef } from 'react';
import FeatherIcon from 'feather-icons-react'

const IndeterminateCheckbox = ({ checked, indeterminate, onChange }) => {
    const checkboxRef = useRef(null);

    useEffect(() => {
        if (checkboxRef.current) {
            checkboxRef.current.indeterminate = indeterminate;
        }        
    }, [indeterminate]);

    return (
        <div className='flex relative'>
            <input
                type="checkbox"
                ref={checkboxRef}
                checked={checked}
                onChange={onChange}
                className='cursor-pointer w-5 h-5 peer appearance-none
                        bg-white border border-subtitle rounded-[4px]
                        indeterminate:bg-primary indeterminate:border-primary
                        checked:bg-primary checked:border-primary
                        '
            />
            <FeatherIcon icon="check" className=" absolute top-[1px] left-[1px] w-[18px] h-[18px] text-white hidden peer-checked:block pointer-events-none" strokeWidth={3}/>
            <FeatherIcon icon="minus" className=" absolute top-[1px] left-[1px] w-[18px] h-[18px] text-white hidden peer-indeterminate:block pointer-events-none" strokeWidth={3}/>
        </div>
    );

}

export default IndeterminateCheckbox