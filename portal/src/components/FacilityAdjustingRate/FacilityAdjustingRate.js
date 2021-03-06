import React, {useEffect, useState} from "react";
import styles from "./FacilityAdjustingRate.module.css";
import {CheckboxItem, FacilityFilterTab, RadioItem} from "../FacilityFilterTab";
import {API_URL} from "../../utils/constants";
import {useAxios} from "../../utils/useAxios";


function FacilityAdjustingRate({
                                   facilities, setFacilities, selectedState, onStateSelected, districtList, selectedDistrict,
                                   setSelectedDistrict, stateList, programs, selectedProgram, setSelectedProgram, facilityType, setFacilityType,
                                   setStatus, fetchFacilities
                               }) {
    const [lastAdjustedOn, setLastAdjustedOn] = useState("Week");
    const [rateWiseFacilities, setRateWiseFacilities] = useState({});
    const [allChecked, setAllChecked] = useState(false);
    const axiosInstance = useAxios('');

    useEffect(() => {
        setStatus("")
    }, []);

    useEffect(() => {
        groupFacilityByRate();
    }, [facilities]);


    const handleChange = (value, setValue) => {
        setValue(value);
    };

    const updateFacility = (index, key, value) => {
        const facilityData = [...facilities];
        facilityData[index][key] = value;
        setFacilities(facilityData);

    };

    const getFacilityProgram = (facility) => {
        if ("programs" in facility) {
            const program = facility.programs.find(obj => obj.id === selectedProgram);
            if (program) {
                return program;
            }
        }
        return {rate: 0};
    };

    const getFacilityList = () => {
        return facilities.map((facility, index) => (
            <tr>
                <td>{facility['facilityCode']}</td>
                <td>{facility['facilityName']}</td>
                <td>{facility['category']}</td>
                <td>{getFacilityProgram(facility).rate}</td>
                <td>
                    <CheckboxItem
                        text={facility['id']}
                        showText={false}
                        checked={facility.isChecked}
                        onSelect={() => {
                            updateFacility(index, "isChecked", !facility.isChecked)
                        }}
                    />

                </td>
            </tr>
        ));

    };

    const handleAllCheck = (e) => {
        let list = [...facilities];
        setAllChecked(e.target.checked);
        list = list.map((ele) => ({
            ...ele,
            isChecked: e.target.checked
        }));
        setFacilities(list);
    };


    const updateRateWiseFacility = (currentRate, newRate) => {
        const rateWiseData = {...rateWiseFacilities};
        rateWiseData[currentRate].newRate = newRate;
        setRateWiseFacilities(rateWiseData);
    };
    const groupFacilityByRate = () => {
        const selectedFacilities = facilities.filter(facility => facility.isChecked);
        const rateWiseData = {};
        selectedFacilities.forEach(facility => {
            const facilityProgram = getFacilityProgram(facility);
            if ("rate" in facilityProgram && facilityProgram.rate in rateWiseData) {
                rateWiseData[facilityProgram.rate].count++;
                rateWiseData[facilityProgram.rate].facilityIds.push(facility.osid);
            } else {
                rateWiseData[facilityProgram.rate || 0] = {
                    count: 1,
                    facilityIds: [facility.osid],
                    newRate: facilityProgram.rate
                }
            }

        });
        setRateWiseFacilities(rateWiseData)
    };

    const getRatesData = () => {
        return Object.keys(rateWiseFacilities).map((rate) => {
            const rateWiseFacilityElement = rateWiseFacilities[rate];
            const facilityCount = rateWiseFacilityElement.count;
            return (<tr>
                <td className="p-1">
                    <RadioItem
                        text={rate}
                        checked={true}
                        onSelect={() => {
                        }}
                        showText={false}
                    />
                </td>
                <td className="p-1">{facilityCount}</td>
                <td className="p-1">{rate}</td>
                <td className="p-1"><input type="number" style={{width: "60px"}} size="4"
                                           value={rateWiseFacilityElement.newRate}
                                           onChange={(evt) => updateRateWiseFacility(parseInt(rate), parseInt(evt.target.value))}/>
                </td>
            </tr>);
        });
    };

    const onSubmitBtnClick = () => {
        setAllChecked(false);
        if (selectedProgram && Object.keys(rateWiseFacilities).length > 0) {
            let updateFacilities = [];
            Object.keys(rateWiseFacilities).forEach((rate) => {
                const facilityIds = rateWiseFacilities[rate].facilityIds;
                facilityIds.forEach((facilityId) => {
                    const facility = facilities.find(facility => facility.osid === facilityId);
                    if (facility) {
                        let programs = [];
                        const program = facility.programs.find(program => program.id === selectedProgram);
                        if (program) {
                            programs = facility.programs.map(program => {
                                if (program.id === selectedProgram) {
                                    return {...program, rate: rateWiseFacilities[rate].newRate};
                                } else {
                                    return program;
                                }
                            })
                        } else {
                            programs = facility.programs.concat({
                                id: selectedProgram,
                                status: "ACTIVE",
                                rate: rateWiseFacilities[rate].newRate
                            })
                        }
                        updateFacilities.push({osid: facility.osid, programs})
                    }
                });


            });
            axiosInstance.current.put(API_URL.FACILITY_API, updateFacilities)
                .then(res => {
                    //registry update in ES happening async, so calling search immediately will not get back actual data
                    setTimeout(() => fetchFacilities(), 1000)
                });
        }
    };


    return (
        <div className={`row ${styles['container']}`}>
            <div className="col-sm-3">
                <FacilityFilterTab
                    programs={programs}
                    setSelectedProgram={setSelectedProgram}
                    states={stateList}
                    setSelectedState={onStateSelected}
                    selectedState={selectedState}
                    districtList={districtList}
                    selectedDistrict={selectedDistrict}
                    setSelectedDistrict={setSelectedDistrict}
                    facilityType={facilityType}
                    setFacilityType={setFacilityType}
                >
                    <div>
                        <span className={"filter-header"}>Last Adjusted on</span>
                        <div className="m-3">
                            <RadioItem
                                text={"Week"}
                                checked={lastAdjustedOn === "Week"}
                                onSelect={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setLastAdjustedOn
                                    )
                                }
                            />
                            <RadioItem
                                text={"Month"}
                                checked={lastAdjustedOn === "Month"}
                                onSelect={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setLastAdjustedOn
                                    )
                                }
                            />
                        </div>
                        <div>
                            <span className={"filter-header"}>Please select date range</span>
                            <div className="m-3">
                                <input className={styles["custom-date-range"]} type="date"/>
                            </div>
                        </div>
                    </div>
                </FacilityFilterTab>
            </div>
            <div className={`col-sm-6 container ${styles['table']}`}>
                <p className={styles['highlight']}>{selectedDistrict} facilties</p>
                <table className={`table table-hover ${styles['table-data']}`}>
                    <thead>
                    <tr>
                        <th>CODE</th>
                        <th>NAME</th>
                        <th>TYPE</th>
                        <th>PROGRAM RATE</th>
                        <th>
                            <CheckboxItem
                                text={"checkAll"}
                                checked={allChecked}
                                onSelect={(e) => {
                                    handleAllCheck(e)
                                }}
                                showText={false}
                            />
                        </th>
                    </tr>
                    </thead>
                    <tbody>{getFacilityList()}</tbody>
                </table>
            </div>
            <div className="col-sm-3 container">
                <div className={styles['highlight']}>Set Rate</div>
                {selectedProgram && Object.keys(rateWiseFacilities).length > 0 && <div>
                    <div
                        className={`overflow-auto text-center table-responsive  ${styles["highlight"]} ${styles["district-table"]}`}>
                        <table className="table table-borderless table-hover">
                            <thead>
                            <tr>
                                <td className="p-1"/>
                                <td className="p-1">No. Of Facilities</td>
                                <td className="p-1">Current Rate</td>
                                <td className="p-1">Set New Rate</td>
                            </tr>
                            </thead>
                            <tbody>{getRatesData()}</tbody>
                        </table>

                    </div>
                    <div className="d-flex justify-content-center">
                        <button className={`${styles['button']} p-2 pl-3 pr-3`} onClick={() => onSubmitBtnClick()}>SET
                            RATES
                        </button>
                    </div>
                    {/*{submit ? <div>All rates set successfully</div> : ''}*/}
                </div>}
            </div>
        </div>
    );
}

export default FacilityAdjustingRate;
