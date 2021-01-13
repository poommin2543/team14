//style
import React, {  useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, Select, InputLabel, MenuItem, TextField, Button, InputAdornment } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Alert, AlertTitle } from '@material-ui/lab';

//api
import { DefaultApi } from '../../api/apis';

//Entity
import { EntPersonal, EntGender, EntTitle, EntDepartment} from '../../api';

//icon
import EmailTwoToneIcon from '@material-ui/icons/EmailTwoTone';
import PersonOutlineTwoToneIcon from '@material-ui/icons/PersonOutlineTwoTone';
import VpnKeyTwoToneIcon from '@material-ui/icons/VpnKeyTwoTone';
import AddCircleTwoToneIcon from '@material-ui/icons/AddCircleTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  margin: {
    margin: theme.spacing(2),
  },
  insideLabel: {
   margin: theme.spacing(1),
  },
  button: {
   marginLeft: '40px',
  },
  textField: {
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
  },
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  pa: {
    marginTop: theme.spacing(2),
  },
  }),
);

export default function Personalpage() {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);

  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);

  const [department, setDepartment] = useState(Number);
  const [title, setTitle] = useState(Number);
  const [gender, setGender] = useState(Number);

  const [personalname, setPersonalname] = useState(String);
  const [email, setEmail] = useState(String);
  const [password, setPassword] = useState(String);

  useEffect(() => {
    const getDepartments = async () => {
      const res = await http.listDepartment({ limit: 10, offset: 0 });
      setLoading(false);
      setDepartments(res);
      console.log(res);
    };
    getDepartments();

    const getTitles = async () => {
      const res = await http.listTitle({ limit: 10, offset: 0 });
      setLoading(false);
      setTitles(res);
      console.log(res);
    };
    getTitles();

    const getGenders = async () => {
      const res = await http.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(res);
      console.log(res);
    };
    getGenders();

    const getPersonals = async () => {
      const res = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();

  }, [loading]);

  const handlePersonalnameChange = (event: any) => {
    setPersonalname(event.target.value as string);
  };

  const handleEmailChange = (event: any) => {
    setEmail(event.target.value as string);
  };

  const handlePasswordChange = (event: any) => {
    setPassword(event.target.value as string);
  };
  
  const TitlehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTitle(event.target.value as number);
  };

  const DepartmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDepartment(event.target.value as number);
  };

  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGender(event.target.value as number);
  };

  // create personal
  const CreatePersonal = async () => {
    if ((personalname != null) && (personalname != "") && (email != null) && (email != "") && (password != null) && (password != "") && (title != null) && (department != null) && (gender != null)){
    const personal = {
      personalname : personalname,
      email : email,
      password : password,
      title : title,
      department : department,
      gender : gender,
    };
    console.log(personals);
    const res: any = await http.createPersonal({ personal: personal });
    console.log("hi");
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      } 
    }
      else {
        setStatus(true);
        setAlert(false);
      }
  };

return (
  <Page theme={pageTheme.tool}>
    <Header
      title="ระบบข้อมูลบุคลากร" type="ระบบแจ้งซ่อมคอมพิวเตอร์">
      <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
        > 
        ออกจากระบบ 
        </Button>
    </Header>
    <Content>
      <ContentHeader title="เพิ่มข้อมูลบุคลากร">
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button  
        onClick={() => {CreatePersonal();}}
        href="/Tablepersonal"
        variant="contained" 
        color="secondary" 
        startIcon={<AddCircleTwoToneIcon/>}
        > 
        บันทึกข้อมูล 
        </Button>
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        component={RouterLink} 
        to="/Tablepersonal" 
        variant="contained"
        color="primary"
        startIcon={<CancelTwoToneIcon/>}
        > 
        ย้อนกลับ 
        </Button>
      </ContentHeader>
      {status ? ( 
        <div>
          {alert ? ( 
              <Alert severity="success"  onClose={() => { }}> 
                <AlertTitle> บันทึกข้อมูลสำเร็จ </AlertTitle></Alert>) 
        : (     
          <Alert severity="error" onClose={() => { setStatus(false) }}> 
            <AlertTitle> ไม่สามารถบันทึกข้อมูลได้ กรุณาลองใหม่อีกครั้ง </AlertTitle></Alert>)}
        </div>
          ) : null}
      <div className={classes.root}>
        <form noValidate autoComplete="off">
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
            <div className={classes.paper}><strong>คำนำหน้าชื่อ</strong></div>
              <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              labelId="nametitle-label"
              id="nametitle"
              value={title}
              onChange={TitlehandleChange}
            >
              <InputLabel className={classes.insideLabel} id="faculty-label">เลือกคำนำหน้าชื่อ</InputLabel>

              {titles.map((item: EntTitle) => (
                <MenuItem value={item.id}>{item.titlename}</MenuItem>
              ))}
            </Select>
            
            <div className={classes.paper}><strong>ชื่อ-นามสกุล</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <PersonOutlineTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="personalname"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={personalname}
              onChange={handlePersonalnameChange}
            />

            <div className={classes.paper}><strong>อีเมลล์</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <EmailTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="Email"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={email}
              onChange={handleEmailChange}
            />

            <div className={classes.paper}><strong>รหัสผ่าน</strong></div>
            <TextField
            style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
            className={classes.textField}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <VpnKeyTwoToneIcon />
                </InputAdornment>
              ),
            }}
              id="Password"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={password}
              onChange={handlePasswordChange}
            />
            
            <div className={classes.paper}><strong>เพศ</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="gender"
              value={gender}
              onChange={GenderhandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกเพศ</InputLabel>

              {genders.map((item: EntGender) => (
                <MenuItem value={item.id}>{item.gendername}</MenuItem>
              ))}
            </Select>

            <div className={classes.paper}><strong>แผนก</strong></div>
            <Select className={classes.select}
              style={{ width: 500 ,marginLeft:7,marginRight:-7,marginTop:10}}
              color="primary"
              id="department"
              value={department}
              onChange={DepartmenthandleChange}
            >
              <InputLabel className={classes.insideLabel}>เลือกแผนกt</InputLabel>

              {departments.map((item: EntDepartment) => (
                <MenuItem value={item.id}>{item.departmentname}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </form>
      </div>
    </Content>
   </Page>
 );
}
