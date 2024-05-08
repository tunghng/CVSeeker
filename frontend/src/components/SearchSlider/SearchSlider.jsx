
import { useState } from "react";
import { Tooltip } from "react-tooltip";

import './SearchSlider.css';

const SearchSlider = () => {
    // ====== State Management ======
    const [sliderValue, setSliderValue] = useState(0.5);
    const valueScripts = {
        0: "Search simply using keywords",
        0.25: "Searching mainly using keywords",
        0.5: "Search with a mix of keywords and context",
        0.75: "Search mainly using context",
        1: "Search based mainly on understanding the context"
    }

    return (
        <div className="flex justify-end items-center gap-3">
            <span
                data-tooltip-id="keyword-tooltip"
                data-tooltip-content="Keyword-based search: Users input keywords to find information"
                data-tooltip-place="bottom"
                data-tooltip-delay-show={200}
                className="cursor-default"
            >
                Keyword
            </span>
            <Tooltip id="keyword-tooltip" />

            <input
                type="range"
                className="cursor-pointer w-40 h-2 appearance-none bg-secondary-subtle hover:bg-primary-subtle rounded-full transition-all duration-300 ease-in-out"
                min="0"
                max="1"
                step="0.25"
                value={sliderValue}
                onChange={(e) => setSliderValue(e.target.value)}
                data-tooltip-id="slider-tooltip"
                data-tooltip-content={valueScripts[sliderValue]}
                data-tooltip-place="bottom"
            />
            <Tooltip id="slider-tooltip" />


            <span
                data-tooltip-id="context-tooltip"
                data-tooltip-content="Context-based search: The search method considers the context or situation to find relevant information"
                data-tooltip-place="bottom"
                data-tooltip-delay-show={200}
                className="cursor-default"
            >
                Context
            </span>
            <Tooltip id="context-tooltip" />
        </div>
    );
};

export default SearchSlider;
