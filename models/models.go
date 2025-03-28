package models

type InspectionForm struct {
	PDF_Path                                                                  string `json:"pdf_path"`
	Bill_To                                                                   string `json:"BILL TO"`
	Location                                                                  string `json:"LOCATION"`
	Bill_To_LN_2                                                              string `json:"BILL TO LN2"`
	Location_LN_2                                                             string `json:"LOCATION LN2"`
	Attention                                                                 string `json:"ATTN"`
	Billing_Street                                                            string `json:"STREET"`
	Billing_Street_LN_2                                                       string `json:"STREET LN2"`
	Location_Street                                                           string `json:"STREET_2"`
	Location_Street_LN_2                                                      string `json:"STREET_2 LN 2"`
	Billing_City_State                                                        string `json:"CITY  STATE"`
	Billing_City_State_LN_2                                                   string `json:"CITY STATE LN 2"`
	Location_City_State                                                       string `json:"CITY  STATE_2"`
	Location_City_State_LN_2                                                  string `json:"CITY STATE_2 LN 2"`
	Contact                                                                   string `json:"CONTACT"`
	Date                                                                      string `json:"DATE"`
	Phone                                                                     string `json:"PHONE"`
	Inspector                                                                 string `json:"INSPECTOR"`
	Email                                                                     string `json:"EMAIL"`
	Inspection_Frequency                                                      string `json:"INSP_FREQ"`
	Inspection_Number                                                         string `json:"INSP_#"`
	Is_The_Building_Occupied                                                  string `json:"1A"`
	Are_All_Systems_In_Service                                                string `json:"1B"`
	Are_FP_Systems_Same_as_Last_Inspection                                    string `json:"1C"`
	Hydraulic_Nameplate_Securely_Attached_And_Legible                         string `json:"1D"`
	Was_A_Main_Drain_Water_Flow_Test_Conducted                                string `json:"2A"`
	Are_All_Sprinkler_System_Main_Control_Valves_Open                         string `json:"3A"`
	Are_All_Other_Valves_In_Proper_Position                                   string `json:"3B"`
	Are_All_Control_Valves_Sealed_Or_Supervised                               string `json:"3C"`
	Are_All_Control_Valves_In_Good_Condition_and_Free_Of_Leaks                string `json:"3D"`
	Are_Fire_Department_Connections_In_Satisfactory_Condition                 string `json:"4A"`
	Are_Caps_In_Place                                                         string `json:"4B"`
	Is_Fire_Department_Connection_Easily_Accessible                           string `json:"4C"`
	Automatic_Drain_Vale_In_Place                                             string `json:"4D"`
	Is_The_Pump_Room_Heated                                                   string `json:"5A"`
	Is_The_Fire_Pump_In_Service                                               string `json:"5B"`
	Was_Fire_Pump_Run_During_This_Inspection                                  string `json:"5C"`
	Was_The_Pump_Started_In_The_Automatic_Mode_By_A_Pressure_Drop             string `json:"5D"`
	Were_The_Pump_Bearings_Lubricated                                         string `json:"5E"`
	Jockey_Pump_Start_Pressure_PSI                                            string `json:"5FPSI"`
	Jockey_Pump_Start_Pressure                                                string `json:"5F"`
	Jockey_Pump_Stop_Pressure_PSI                                             string `json:"5GPSI"`
	Jockey_Pump_Stop_Pressure                                                 string `json:"5G"`
	Fire_Pump_Start_Pressure_PSI                                              string `json:"5HPSI"`
	Fire_Pump_Start_Pressure                                                  string `json:"5H"`
	Fire_Pump_Stop_Pressure_PSI                                               string `json:"5IPSI"`
	Fire_Pump_Stop_Pressure                                                   string `json:"5I"`
	Is_The_Fuel_Tank_At_Least_2_3_Full                                        string `json:"6A"`
	Is_Engine_Oil_At_Correct_Level                                            string `json:"6B"`
	Is_Engine_Coolant_At_Correct_Level                                        string `json:"6C"`
	Is_The_Engine_Block_Heater_Working                                        string `json:"6D"`
	Is_Pump_Room_Ventilation_Operating_Properly                               string `json:"6E"`
	Was_Water_Discharge_Observed_From_Heat_Exchanger_Return_Line              string `json:"6F"`
	Was_Cooling_Line_Strainer_Cleaned_After_Test                              string `json:"6G"`
	Was_Pump_Run_For_At_Least_30_Minutes                                      string `json:"6H"`
	Does_The_Switch_In_Auto_Alarm_Work                                        string `json:"6I"`
	Does_The_Pump_Running_Alarm_Work                                          string `json:"6J"`
	Does_The_Common_Alarm_Work                                                string `json:"6K"`
	Was_Casing_Relief_Valve_Operating_Properly                                string `json:"7A"`
	Was_Pump_Run_For_At_Least_10_Minutes                                      string `json:"7B"`
	Does_The_Loss_Of_Power_Alarm_Work                                         string `json:"7C"`
	Does_The_Electric_Pump_Running_Alarm_Work                                 string `json:"7D"`
	Power_Failure_Condition_Simulated_While_Pump_Operating_At_Peak_Load       string `json:"7E"`
	Trasfer_Of_Power_To_Alternative_Power_Source_Verified                     string `json:"7F"`
	Power_Faulure_Condition_Removed                                           string `json:"7G"`
	Pump_Reconnected_To_Normal_Power_Source_After_A_Time_Delay                string `json:"7H"`
	Have_Anti_Freeze_Systems_Been_Tested                                      string `json:"8A"`
	Freeze_Protection_In_Degrees_F                                            string `json:"AFTEMP"`
	Are_Alarm_Valves_Water_Flow_Devices_And_Retards_In_Satisfactory_Condition string `json:"8B"`
	Water_Flow_Alarm_Test_Conducted_With_Inspectors_Test                      string `json:"8C"`
	Water_Flow_Alarm_Test_Conducted_With_Bypass_Connection                    string `json:"8D"`
	Is_Dry_Valve_In_Service_And_In_Good_Condition                             string `json:"9A"`
	Is_Dry_Valve_Itermediate_Chamber_Not_Leaking                              string `json:"9B"`
	Has_The_Dry_System_Been_Fully_Tripped_Within_The_Last_Three_Years         string `json:"9C"`
	Are_Quick_Opening_Device_Control_Valves_Open                              string `json:"9D"`
	Is_There_A_List_Of_Known_Low_Point_Drains_At_The_Riser                    string `json:"9E"`
	Have_Known_Low_Points_Been_Drained                                        string `json:"9F"`
	Is_Oil_Level_Full_On_Air_Compressor                                       string `json:"9G"`
	Does_The_Air_Compressor_Return_System_Pressure_In_30_minutes_or_Under     string `json:"9H"`
	What_Pressure_Does_Air_Compressor_Start_PSI                               string `json:"9ISTARTPSI"`
	What_Pressure_Does_Air_Compressor_Start                                   string `json:"9I"`
	What_Pressure_Does_Air_Compressor_Stop_PSI                                string `json:"9JPSISTOP"`
	What_Pressure_Does_Air_Compressor_Stop                                    string `json:"9J"`
	Did_Low_Air_Alarm_Operate_PSI                                             string `json:"9KLOWAIR"`
	Did_Low_Air_Alarm_Operate                                                 string `json:"9K"`
	Date_Of_Last_Full_Trip_Test                                               string `json:"LASTFULLTRIP"`
	Date_Of_Last_Internal_Inspection                                          string `json:"LASTINTERNAL"`
	Are_Valves_In_Service_And_In_Good_Condition                               string `json:"10a"`
	Were_Valves_Tripped                                                       string `json:"10b"`
	What_Pressure_Did_Pneumatic_Actuator_Trip_PSI                             string `json:"10c PSI"`
	What_Pressure_Did_Pneumatic_Actuator_Trip                                 string `json:"10c"`
	Was_Priming_Line_Left_On_After_Test                                       string `json:"10d"`
	What_Pressure_Does_Preaction_Air_Compressor_Start_PSI                     string `json:"10e PSI"`
	What_Pressure_Does_Preaction_Air_Compressor_Start                         string `json:"10e"`
	What_Pressure_Does_Preaction_Air_Compressor_Stop_PSI                      string `json:"10f PSI"`
	What_Pressure_Does_Preaction_Air_Compressor_Stop                          string `json:"10f"`
	Did_Preaction_Low_Air_Alarm_Operate_PSI                                   string `json:"10g PSI"`
	Did_Preaction_Low_Air_Alarm_Operate                                       string `json:"10g"`
	Does_Water_Motor_Gong_Work                                                string `json:"11a"`
	Does_Electric_Bell_Work                                                   string `json:"11b"`
	Are_Water_Flow_Alarms_Operational                                         string `json:"11c"`
	Are_All_Tamper_Switches_Operational                                       string `json:"11d"`
	Did_Alarm_Panel_Clear_After_Test                                          string `json:"11e"`
	Are_A_Minimum_Of_6_Spare_Sprinklers_Readily_Avaiable                      string `json:"12a"`
	Is_Condition_Of_Piping_And_Other_System_Componets_Satisfactory            string `json:"12b"`
	Are_Known_Dry_Type_Heads_Less_Than_10_Years_Old                           string `json:"12c"`
	Are_Known_Dry_Type_Heads_Less_Than_10_Years_Old_Year                      string `json:"12c year"`
	Are_Known_Quick_Response_Heads_Less_Than_20_Years_Old                     string `json:"12d"`
	Are_Known_Quick_Response_Heads_Less_Than_20_Years_Old_Year                string `json:"12d year"`
	Are_Known_Standard_Response_Heads_Less_Than_50_Years_Old                  string `json:"12e"`
	Are_Known_Standard_Response_Heads_Less_Than_50_Years_Old_Year             string `json:"12e year"`
	Have_All_Gauges_Been_Tested_Or_Replaced_In_The_Last_5_Years               string `json:"12f"`
	Have_All_Gauges_Been_Tested_Or_Replaced_In_The_Last_5_Years_Year          string `json:"12f year"`
	System_1_Name                                                             string `json:"Drain test line 1"`
	System_1_Drain_Size                                                       string `json:"drian size 1"`
	System_1_Static_PSI                                                       string `json:"Static 1"`
	System_1_Residual_PSI                                                     string `json:"Residual 1"`
	System_2_Name                                                             string `json:"Drain test line 2"`
	System_2_Drain_Size                                                       string `json:"drian size 2"`
	System_2_Static_PSI                                                       string `json:"Static 2"`
	System_2_Residual_PSI                                                     string `json:"Residual 2"`
	System_3_Name                                                             string `json:"Drain test line 3"`
	System_3_Drain_Size                                                       string `json:"drian size 3"`
	System_3_Static_PSI                                                       string `json:"Static 3"`
	System_3_Residual_PSI                                                     string `json:"Residual 3"`
	System_4_Name                                                             string `json:"Drain test line 4"`
	System_4_Drain_Size                                                       string `json:"drian size 4"`
	System_4_Static_PSI                                                       string `json:"Static 4"`
	System_4_Residual_PSI                                                     string `json:"Residual 4"`
	System_5_Name                                                             string `json:"Drain test line 5"`
	System_5_Drain_Size                                                       string `json:"drian size 5"`
	System_5_Static_PSI                                                       string `json:"Static 5"`
	System_5_Residual_PSI                                                     string `json:"Residual 5"`
	System_6_Name                                                             string `json:"Drain test line 6"`
	System_6_Drain_Size                                                       string `json:"drian size 6"`
	System_6_Static_PSI                                                       string `json:"Static 6"`
	System_6_Residual_PSI                                                     string `json:"Residual 6"`
	Drain_Test_Notes                                                          string `json:"Drain test notes"`
	Device_1_Name                                                             string `json:"15a pt1"`
	Device_1_Address                                                          string `json:"15a pt2"`
	Device_1_Description_Location                                             string `json:"15a pt3"`
	Device_1_Operated                                                         string `json:"15a pt4"`
	Device_1_Delay_Sec                                                        string `json:"15a pt5"`
	Device_2_Name                                                             string `json:"15b pt1"`
	Device_2_Address                                                          string `json:"15b pt2"`
	Device_2_Description_Location                                             string `json:"15b pt3"`
	Device_2_Operated                                                         string `json:"15b pt4"`
	Device_2_Delay_Sec                                                        string `json:"15b pt5"`
	Device_3_Name                                                             string `json:"15c pt1"`
	Device_3_Address                                                          string `json:"15c pt2"`
	Device_3_Description_Location                                             string `json:"15c pt3"`
	Device_3_Operated                                                         string `json:"15c pt4"`
	Device_3_Delay_Sec                                                        string `json:"15c pt5"`
	Device_4_Name                                                             string `json:"15d pt1"`
	Device_4_Address                                                          string `json:"15d pt2"`
	Device_4_Description_Location                                             string `json:"15d pt3"`
	Device_4_Operated                                                         string `json:"15d pt4"`
	Device_4_Delay_Sec                                                        string `json:"15d pt5"`
	Device_5_Name                                                             string `json:"15e pt1"`
	Device_5_Address                                                          string `json:"15e pt2"`
	Device_5_Description_Location                                             string `json:"15e pt3"`
	Device_5_Operated                                                         string `json:"15e pt4"`
	Device_5_Delay_Sec                                                        string `json:"15e pt5"`
	Device_6_Name                                                             string `json:"15f pt1"`
	Device_6_Address                                                          string `json:"15f pt2"`
	Device_6_Description_Location                                             string `json:"15f pt3"`
	Device_6_Operated                                                         string `json:"15f pt4"`
	Device_6_Delay_Sec                                                        string `json:"15f pt5"`
	Device_7_Name                                                             string `json:"15g pt1"`
	Device_7_Address                                                          string `json:"15g pt2"`
	Device_7_Description_Location                                             string `json:"15g pt3"`
	Device_7_Operated                                                         string `json:"15g pt4"`
	Device_7_Delay_Sec                                                        string `json:"15g pt5"`
	Device_8_Name                                                             string `json:"15h pt1"`
	Device_8_Address                                                          string `json:"15h pt2"`
	Device_8_Description_Location                                             string `json:"15h pt3"`
	Device_8_Operated                                                         string `json:"15h pt4"`
	Device_8_Delay_Sec                                                        string `json:"15h pt5"`
	Device_9_Name                                                             string `json:"15i pt1"`
	Device_9_Address                                                          string `json:"15i pt2"`
	Device_9_Description_Location                                             string `json:"15i pt3"`
	Device_9_Operated                                                         string `json:"15i pt4"`
	Device_9_Delay_Sec                                                        string `json:"15i pt5"`
	Device_10_Name                                                            string `json:"15j pt1"`
	Device_10_Address                                                         string `json:"15j pt2"`
	Device_10_Description_Location                                            string `json:"15j pt3"`
	Device_10_Operated                                                        string `json:"15j pt4"`
	Device_10_Delay_Sec                                                       string `json:"15j pt5"`
	Device_11_Name                                                            string `json:"15k pt1"`
	Device_11_Address                                                         string `json:"15k pt2"`
	Device_11_Description_Location                                            string `json:"15k pt3"`
	Device_11_Operated                                                        string `json:"15k pt4"`
	Device_11_Delay_Sec                                                       string `json:"15k pt5"`
	Device_12_Name                                                            string `json:"15l pt1"`
	Device_12_Address                                                         string `json:"15l pt2"`
	Device_12_Description_Location                                            string `json:"15l pt3"`
	Device_12_Operated                                                        string `json:"15l pt4"`
	Device_12_Delay_Sec                                                       string `json:"15l pt5"`
	Device_13_Name                                                            string `json:"15m pt1"`
	Device_13_Address                                                         string `json:"15m pt2"`
	Device_13_Description_Location                                            string `json:"15m pt3"`
	Device_13_Operated                                                        string `json:"15m pt4"`
	Device_13_Delay_Sec                                                       string `json:"15m pt5"`
	Device_14_Name                                                            string `json:"15n pt1"`
	Device_14_Address                                                         string `json:"15n pt2"`
	Device_14_Description_Location                                            string `json:"15n pt3"`
	Device_14_Operated                                                        string `json:"15n pt4"`
	Device_14_Delay_Sec                                                       string `json:"15n pt5"`
	Adjustments_Or_Corrections_Make                                           string `json:"16 Adjustments or Corrections"`
	Explanation_Of_Any_No_Answers                                             string `json:"17 Explanation of no answers"`
	Explanation_Of_Any_No_Answers_Continued                                   string `json:"17 Explanation of no answers continued"`
	Notes                                                                     string `json:"18 NOTES"`
}

type DryForm struct {
	PDF_Path                                        string `json:"pdf_path"`
	Report_To                                       string `json:"REPORT TO"`
	Building                                        string `json:"Building,alt=Text1"`
	Report_To_2                                     string `json:"Report to cont,alt=Text6"`
	Building_2                                      string `json:"Building cont,alt=Text5"`
	Attention                                       string `json:"Attn,alt=Text2"`
	Street                                          string `json:"STREET"`
	Inspector                                       string `json:"INSPECTOR,alt=Dropdown1"`
	City_State                                      string `json:"CITY  STATE"`
	Date                                            string `json:"DATE"`
	Dry_Pipe_Valve_Make                             string `json:"MAKE"`
	Dry_Pipe_Valve_Model                            string `json:"MODEL"`
	Dry_Pipe_Valve_Size                             string `json:"SIZE"`
	Dry_Pipe_Valve_Year                             string `json:"YEAR"`
	Dry_Pipe_Valve_Controls_Sprinklers_In           string `json:"CONTROLS SPRINKLERS IN"`
	Quick_Opening_Device_Make                       string `json:"MAKE_2"`
	Quick_Opening_Device_Model                      string `json:"MODEL_2"`
	Quick_Opening_Device_Control_Valve_Open         string `json:"CONTROL VALVE OPEN"`
	Quick_Opening_Device_Year                       string `json:"YEAR_2"`
	Trip_Test_Air_Pressure_Before_Test              string `json:"AIR PSI"`
	Trip_Test_Air_System_Tripped_At                 string `json:"AIR PSI2"`
	Trip_Test_Water_Pressure_Before_Test            string `json:"WATER PSI"`
	Trip_Test_Time                                  string `json:"Time,alt=Text3"`
	Trip_Test_Air_Quick_Opening_Device_Operated_At  string `json:"AIR PSIQOD OPERATED AT"`
	Trip_Test_Time_Quick_Opening_Device_Operated_At string `json:"TIMEQOD OPERATED AT"`
	Trip_Test_Time_Water_At_Inspectors_Test         string `json:"WATER AT INSPECTORS TEST"`
	Trip_Test_Static_Water_Pressure                 string `json:"PSI"`
	Trip_Test_Residual_Water_Pressure               string `json:"PSI_2"`
	Remarks_On_Test                                 string `json:"Remarks,alt=Text4"`
}

type PumpForm struct {
	PDF_Path                      string `json:"pdf_path"`
	Report_To                     string `json:"REPORT TO"`
	Building                      string `json:"BUILDINGATTN"`
	Attention                     string `json:"ATTN:"`
	Street                        string `json:"STREET:"`
	Inspector                     string `json:"INSPECTORSELECT:"`
	City_State                    string `json:"CITY  STATE:"`
	Date                          string `json:"DATE"`
	Pump_Make                     string `json:"MAKE"`
	Pump_Rated_RPM                string `json:"RATED RPM"`
	Pump_Model                    string `json:"MODEL"`
	Pump_Rated_GPM                string `json:"RATED GPM"`
	Pump_Serial_Number            string `json:"SN"`
	Pump_Max_PSI                  string `json:"SMAX PSI"`
	Pump_Power                    string `json:"Dropdown10"`
	Pump_Rated_PSI                string `json:"RATED PSI"`
	Pump_Water_Supply             string `json:"Dropdown9"`
	Pump_PSI_At_150_Percent       string `json:"PSI  150"`
	Pump_Controller_Make          string `json:"CONTROLERMAKE"`
	Pump_Controller_Voltage       string `json:"VOLTAGE"`
	Pump_Controller_Model         string `json:"CONTROLERMODEL"`
	Pump_Controller_Horse_Power   string `json:"HP"`
	Pump_Controller_Serial_Number string `json:"CONTROLERSN"`
	Pump_Controller_Supervision   string `json:"SUPERVISION"`
	Diesel_Engine_Make            string `json:"DENGMAKE"`
	Diesel_Engine_Serial_Number   string `json:"DENGSN"`
	Diesel_Engine_Model           string `json:"DENGMODEL"`
	Diesel_Engine_Hours           string `json:"DENGHOURS"`
	Flow_Test_Orifice_Size_1      string `json:"Dropdown2"`
	Flow_Test_Orifice_Size_2      string `json:"Dropdown3"`
	Flow_Test_Orifice_Size_3      string `json:"Dropdown4"`
	Flow_Test_Orifice_Size_4      string `json:"Dropdown5"`
	Flow_Test_Orifice_Size_5      string `json:"Dropdown6"`
	Flow_Test_Orifice_Size_6      string `json:"Dropdown7"`
	Flow_Test_Orifice_Size_7      string `json:"Dropdown8"`
	Flow_Test_1_Suction_PSI       string `json:"SUCTIONPSIRow1"`
	Flow_Test_1_Discharge_PSI     string `json:"DISCHARGEPSIRow1"`
	Flow_Test_1_Net_PSI           string `json:"PITOT"`
	Flow_Test_1_RPM               string `json:"RPM1"`
	Flow_Test_1_O1_Pitot          string `json:"1"`
	Flow_Test_1_O2_Pitot          string `json:"2"`
	Flow_Test_1_O3_Pitot          string `json:"3"`
	Flow_Test_1_O4_Pitot          string `json:"4"`
	Flow_Test_1_O5_Pitot          string `json:"5"`
	Flow_Test_1_O6_Pitot          string `json:"6"`
	Flow_Test_1_O7_Pitot          string `json:"7"`
	Flow_Test_1_O1_GPM            string `json:"8"`
	Flow_Test_1_O2_GPM            string `json:"9"`
	Flow_Test_1_O3_GPM            string `json:"10"`
	Flow_Test_1_O4_GPM            string `json:"11"`
	Flow_Test_1_O5_GPM            string `json:"12"`
	Flow_Test_1_O6_GPM            string `json:"13"`
	Flow_Test_1_O7_GPM            string `json:"14"`
	Flow_Test_1_Total_Flow        string `json:"TOTAL FLOW 1"`
	Flow_Test_2_Suction_PSI       string `json:"SUCTIONPSIRow2"`
	Flow_Test_2_Discharge_PSI     string `json:"DISCHARGEPSIRow2"`
	Flow_Test_2_Net_PSI           string `json:"PITOT_2"`
	Flow_Test_2_RPM               string `json:"RPM 2"`
	Flow_Test_2_O1_Pitot          string `json:"15"`
	Flow_Test_2_O2_Pitot          string `json:"16"`
	Flow_Test_2_O3_Pitot          string `json:"17"`
	Flow_Test_2_O4_Pitot          string `json:"18"`
	Flow_Test_2_O5_Pitot          string `json:"19"`
	Flow_Test_2_O6_Pitot          string `json:"20"`
	Flow_Test_2_O7_Pitot          string `json:"21"`
	Flow_Test_2_O1_GPM            string `json:"22"`
	Flow_Test_2_O2_GPM            string `json:"23"`
	Flow_Test_2_O3_GPM            string `json:"24"`
	Flow_Test_2_O4_GPM            string `json:"25"`
	Flow_Test_2_O5_GPM            string `json:"26"`
	Flow_Test_2_O6_GPM            string `json:"27"`
	Flow_Test_2_O7_GPM            string `json:"28"`
	Flow_Test_2_Total_Flow        string `json:"TOTAL FLOW 2"`
	Flow_Test_3_Suction_PSI       string `json:"SUCTIONPSIRow3"`
	Flow_Test_3_Discharge_PSI     string `json:"DISCHARGEPSIRow3"`
	Flow_Test_3_Net_PSI           string `json:"PITOT_3"`
	Flow_Test_3_RPM               string `json:"RPM 3"`
	Flow_Test_3_O1_Pitot          string `json:"29"`
	Flow_Test_3_O2_Pitot          string `json:"30"`
	Flow_Test_3_O3_Pitot          string `json:"31"`
	Flow_Test_3_O4_Pitot          string `json:"32"`
	Flow_Test_3_O5_Pitot          string `json:"33"`
	Flow_Test_3_O6_Pitot          string `json:"34"`
	Flow_Test_3_O7_Pitot          string `json:"35"`
	Flow_Test_3_O1_GPM            string `json:"36"`
	Flow_Test_3_O2_GPM            string `json:"37"`
	Flow_Test_3_O3_GPM            string `json:"38"`
	Flow_Test_3_O4_GPM            string `json:"39"`
	Flow_Test_3_O5_GPM            string `json:"40"`
	Flow_Test_3_O6_GPM            string `json:"41"`
	Flow_Test_3_O7_GPM            string `json:"42"`
	Flow_Test_3_Total_Flow        string `json:"TOTAL FLOW 3"`
	Flow_Test_4_Suction_PSI       string `json:"SUCTIONPSIRow4"`
	Flow_Test_4_Discharge_PSI     string `json:"DISCHARGEPSIRow4"`
	Flow_Test_4_Net_PSI           string `json:"PITOT_4"`
	Flow_Test_4_RPM               string `json:"RPM 4"`
	Flow_Test_4_O1_Pitot          string `json:"43"`
	Flow_Test_4_O2_Pitot          string `json:"44"`
	Flow_Test_4_O3_Pitot          string `json:"45"`
	Flow_Test_4_O4_Pitot          string `json:"46"`
	Flow_Test_4_O5_Pitot          string `json:"47"`
	Flow_Test_4_O6_Pitot          string `json:"48"`
	Flow_Test_4_O7_Pitot          string `json:"49"`
	Flow_Test_4_O1_GPM            string `json:"50"`
	Flow_Test_4_O2_GPM            string `json:"51"`
	Flow_Test_4_O3_GPM            string `json:"52"`
	Flow_Test_4_O4_GPM            string `json:"53"`
	Flow_Test_4_O5_GPM            string `json:"54"`
	Flow_Test_4_O6_GPM            string `json:"55"`
	Flow_Test_4_O7_GPM            string `json:"56"`
	Flow_Test_4_Total_Flow        string `json:"TOTAL FLOW 4"`
	Remarks_On_Test               string `json:"Text11"`
}

type BackflowForm struct {
	PDF_Path                                string `json:"pdf_path"`
	Owner_Of_Property                       string `json:"Bill to,alt=Text2"`
	Date                                    string `json:"Date,alt=Text5"`
	Mailing_Address                         string `json:"Mailing,alt=Text3"`
	Tested_By                               string `json:"Tested by,alt=Dropdown8"`
	Certificate_Number                      string `json:"Certificate #,alt=Dropdown9"`
	Contact_Person                          string `json:"Contact,alt=Text4"`
	Backflow_Type                           string `json:"BF Type,alt=Group11"`
	Backflow_Make                           string `json:"BF Make,alt=Dropdown12"`
	Backflow_Model                          string `json:"BF model,alt=Text7"`
	Backflow_Size                           string `json:"BF size,alt=Dropdown13"`
	Backflow_Serial_Number                  string `json:"BF SN,alt=Text14"`
	Test_Type                               string `json:"Test freq,alt=Group2"`
	Device_Location                         string `json:"Device location,alt=Text1"`
	RPZ_Check_Valve_1_Closed_Tight          string `json:"Check Box3"`
	RPZ_Check_Valve_1_Leaked                string `json:"Check Box4"`
	RPZ_Check_Valve_1_PSID                  string `json:"rpzck 1"`
	RPZ_Check_Valve_2_Closed_Tight          string `json:"Check Box5"`
	RPZ_Check_Valve_Flow                    string `json:"Check Box7"`
	RPZ_Check_Valve_No_Flow                 string `json:"Check Box8"`
	PVB_SRVB_Check_Valve_Flow               string `json:"Group6"`
	RPZ_Relief_Valve_Opened_At_PSID         string `json:"relief psi"`
	RPZ_Check_Valve_2_PSID                  string `json:"rpzck2"`
	RPZ_Check_Valve_2_Leaked                string `json:"Check Box6"`
	PVB_SRVB_Check_Valve_PSID               string `json:"pvb"`
	RPZ_Relief_Valve_Did_Not_Open           string `json:"Check Box15"`
	PVB_SRVB_Air_Inlet_Valve_Opened_At_PSID string `json:"pvb open"`
	DCVA_Back_Pressure_Test_1_PSI           string `json:"tc1 psi"`
	DCVA_Back_Pressure_Test_4_PSI           string `json:"tc2 psi"`
	DCVA_Check_Valve_1_PSID                 string `json:"dc ck1"`
	DCVA_Check_Valve_2_PSID                 string `json:"dc ck2"`
	DCVA_Flow                               string `json:"Check Box9"`
	DCVA_No_Flow                            string `json:"Check Box10"`
	Downsteam_Shutoff_Valve_Status          string `json:"downstream"`
	PVB_SRVB_Air_Inlet_Valve_Did_Not_Open   string `json:"Check Box21"`
	Protection_Type                         string `json:"service type"`
	Result                                  string `json:"pass fail"`
	Remarks_1                               string `json:"bf remark"`
	Remarks_2                               string `json:"bf remark2"`
	Witness                                 string `json:"witness"`
	Remarks_3                               string `json:"bf remark3"`
}
type BackflowForm2 struct {
	PDF_Path                                string `json:"pdf_path"`
	Owner_Of_Property                       string `json:"Text2"`
	Date                                    string `json:"Text5"`
	Mailing_Address                         string `json:"Text3"`
	Tested_By                               string `json:"Dropdown8"`
	Certificate_Number                      string `json:"Dropdown9"`
	Contact_Person                          string `json:"Text4"`
	Backflow_Type                           string `json:"Group11"`
	Backflow_Make                           string `json:"Dropdown12"`
	Backflow_Model                          string `json:"Text7"`
	Backflow_Size                           string `json:"Dropdown13"`
	Backflow_Serial_Number                  string `json:"Text14"`
	Test_Type                               string `json:"Group2"`
	Device_Location                         string `json:"Text1"`
	RPZ_Check_Valve_1_Closed_Tight          string `json:"Check Box3"`
	RPZ_Check_Valve_1_Leaked                string `json:"Check Box4"`
	RPZ_Check_Valve_1_PSID                  string `json:"rpzck 1"`
	RPZ_Check_Valve_2_Closed_Tight          string `json:"Check Box5"`
	RPZ_Check_Valve_Flow                    string `json:"Check Box7"`
	RPZ_Check_Valve_No_Flow                 string `json:"Check Box8"`
	PVB_SRVB_Check_Valve_Flow               string `json:"Group6"`
	RPZ_Relief_Valve_Opened_At_PSID         string `json:"relief psi"`
	RPZ_Check_Valve_2_PSID                  string `json:"rpzck2"`
	RPZ_Check_Valve_2_Leaked                string `json:"Check Box6"`
	PVB_SRVB_Check_Valve_PSID               string `json:"pvb"`
	RPZ_Relief_Valve_Did_Not_Open           string `json:"Check Box15"`
	PVB_SRVB_Air_Inlet_Valve_Opened_At_PSID string `json:"pvb open"`
	DCVA_Back_Pressure_Test_1_PSI           string `json:"tc1 psi"`
	DCVA_Back_Pressure_Test_4_PSI           string `json:"tc2 psi"`
	DCVA_Check_Valve_1_PSID                 string `json:"dc ck1"`
	DCVA_Check_Valve_2_PSID                 string `json:"dc ck2"`
	DCVA_Flow                               string `json:"Check Box9"`
	DCVA_No_Flow                            string `json:"Check Box10"`
	Downsteam_Shutoff_Valve_Status          string `json:"downstream"`
	PVB_SRVB_Air_Inlet_Valve_Did_Not_Open   string `json:"Check Box21"`
	Protection_Type                         string `json:"service type"`
	Result                                  string `json:"pass fail"`
	Remarks_1                               string `json:"bf remark"`
	Remarks_2                               string `json:"bf remark2"`
	Witness                                 string `json:"witness"`
	Remarks_3                               string `json:"bf remark3"`
}
