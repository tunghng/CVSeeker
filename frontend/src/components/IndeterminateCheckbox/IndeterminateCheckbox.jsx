
import { useEffect, useRef } from 'react';

const IndeterminateCheckbox = ({ checked, indeterminate, onChange }) => {
    const checkboxRef = useRef(null);

    useEffect(() => {
        checkboxRef.current.indeterminate = indeterminate;
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
            <svg className='feather feather-check absolute top-[1px] left-[1px] w-[18px] h-[18px] text-white hidden peer-checked:block pointer-events-none' xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
            <svg className='feather feather-minus absolute top-[1px] left-[1px] w-[18px] h-[18px] text-white hidden peer-indeterminate:block pointer-events-none' xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round"><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        </div>
    );

}

export default IndeterminateCheckbox