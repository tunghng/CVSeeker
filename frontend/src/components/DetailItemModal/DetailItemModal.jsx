
import FeatherIcon from 'feather-icons-react'
import './DetailItemModal.css'

const DetailItemModal = ({ showDetailItemModal, detailItem, onModalClose, onAddToList, onDownloadClick }) => {

    if (!detailItem) return null

    return (
        <div className={`${showDetailItemModal ? 'block' : 'hidden'} fixed top-0 left-0 w-screen min-h-screen z-20 bg-black/80`}
            onClick={onModalClose}>

            {/* ====== Modal Top Bar ====== */}
            <div className='fixed w-full h-14 px-6 flex justify-between items-center bg-black/80'
                onClick={(e) => e.stopPropagation()}>
                <div className='flex items-baseline gap-x-6'>
                    <h1 className='text-xl text-white'>{detailItem?.basic_info.full_name}</h1>
                </div>

                <div className='flex gap-x-3'>
                    <button className='px-3 py-2 sm:py-1 bg-white/20 text-base text-white rounded-full flex items-center gap-x-1 hover:bg-white/30'
                        onClick={onDownloadClick}>
                        <FeatherIcon icon="download" className='w-[18px] h-[18px]' strokeWidth={1.6} />
                        <p className='hidden sm:block'>Download</p>
                    </button>
                    {onAddToList && (
                        <button className='px-3 py-2 sm:py-1 bg-white/20 text-base text-white rounded-full flex items-center gap-x-1 hover:bg-white/30'
                            onClick={onAddToList}>
                            <FeatherIcon icon="plus" className='w-[18px] h-[18px]' strokeWidth={1.6} />
                            <p className='hidden sm:block'>Add to List</p>
                        </button>
                    )}
                </div>
            </div>

            {/* ====== Modal Content ====== */}
            <div className='pt-20 pb-16 h-screen overflow-auto'>
                <div className='my-container-medium pb-10 bg-background'
                    onClick={(e) => e.stopPropagation()}>

                    <h1 className='table-section-info'>{detailItem?.basic_info.full_name}</h1>

                    <h2 className='pt-3 pb-4 px-2'>{detailItem?.summary}</h2>

                    <div className='table-section-container'>
                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>Education Level</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <p className='table-section-data'>{detailItem?.basic_info.education_level}</p>
                            </div>
                        </div>

                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>GPA</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <p className='table-section-data'>{detailItem?.basic_info.gpa}</p>
                            </div>
                        </div>

                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>Majors</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <p className='table-section-data'>{detailItem.basic_info.majors.length > 0 ? detailItem.basic_info.majors.join(', ') : 'N/A'}</p>
                            </div>
                        </div>

                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>University</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <p className='table-section-data'>{detailItem?.basic_info.university}</p>
                            </div>
                        </div>
                    </div>


                    <div className='table-section-container mt-5'>
                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>Skills</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <ul>
                                    {detailItem.skills.length > 0 ?
                                        detailItem.skills.map((skill, index) => (
                                            <li key={index} className='table-section-data before:content-["•"] before:mr-2'>
                                                {skill}
                                            </li>
                                        ))
                                        : 'None'}
                                </ul>
                            </div>
                        </div>
                    </div>

                    <div className='table-section-container mt-5'>
                        <div className='table-section-row'>
                            <div className='table-section-wrapper'>
                                <p className='table-section-title'>Awards</p>
                            </div>
                            <div className='table-section-wrapper'>
                                <ul>
                                    {detailItem.award.length > 0 ?
                                        detailItem.award.map((award, index) => (
                                            <li key={index} className='table-section-data before:content-["•"] before:mr-2'>
                                                {award.award_name}
                                            </li>
                                        ))
                                        : 'None'}
                                </ul>
                            </div>
                        </div>
                    </div>

                    <h1 className='table-section-info'>Work experience</h1>
                    {detailItem.work_experience.length > 0
                        ?
                        detailItem.work_experience.map((work, index) => (
                            <div key={index} className='mt-2 px-2 space-y-1'>
                                <div className='flex justify-between'>
                                    <p className='table-section-title text-lg'>{work.job_title}</p>
                                    <p className='table-section-data italic'>{work.duration}</p>
                                </div>
                                <p>Company: {work.company} ({work.location})</p>
                                <p>Summary: {work.job_summary}</p>
                            </div>
                        ))
                        :
                        <p className='mt-2 px-2'>None</p>
                    }


                    <h1 className='table-section-info'>Project experience</h1>
                    {detailItem.project_experience.length > 0
                        ?
                        detailItem.project_experience.map((project, index) => (
                            <div key={index} className='mt-2 px-2 space-y-1'>
                                <p className='table-section-title text-lg'>{project.project_name}</p>
                                <p className='table-section-data'>{project.project_description}</p>
                            </div>
                        ))
                        :
                        <p className='mt-2 px-2'>None</p>
                    }

                </div>
            </div>
        </div>
    )
}

export default DetailItemModal
