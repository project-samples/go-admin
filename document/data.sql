create table code_master (
    master varchar(100) not null,
    code varchar(100) not null,
    name varchar(100),
    sequence int8,
    status char(1),
    primary key (master, code)
)
create table modules (
  moduleid varchar(40) primary key,
  modulename varchar(255) not null,
  status char(1) not null,
  path varchar(255),
  resourcekey varchar(255),
  icon varchar(255),
  sequence int not null,
  actions int4 null,
  parent varchar(40),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz
);

create table users (
  userid varchar(40) primary key,
  username varchar(255) not null,
  email varchar(255) not null,
  displayname varchar(255) not null,
  status char(1) not null,
  gender char(1),
  phone varchar(20),
  title varchar(10),
  position varchar(40),
  imageurl varchar(500),
  language varchar(5),
  dateformat varchar(12),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz,
  lastlogin timestamptz
);

create table roles (
  roleid varchar(40) primary key,
  rolename varchar(255) not null,
  status char(1) not null,
  remark varchar(255),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz
);
create table userroles (
  userid varchar(40) not null,
  roleid varchar(40) not null,
  primary key (userid, roleid)
);
create table rolemodules (
  roleid varchar(40) not null,
  moduleid varchar(40) not null,
  permissions int not null,
  primary key (roleid, moduleid)
);

create table auditlog (
  id varchar(255) primary key,
  resource varchar(255),
  userid varchar(255),
  ip varchar(255),
  action varchar(255),
  time timestamptz,
  status varchar(255),
  remark varchar(255)
);
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('dashboard','Dashboard','A','/dashboard','dashboard','assignments',1,7,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('admin','Admin','A','/admin','admin','contacts',2,7,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('setup','Setup','A','/setup','setup','settings',3,7,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('report','Report','A','/report','report','pie_chart',4,7,'');

insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('user','User Management','A','/users','user','person',1,7,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('role','Role Management','A','/roles','role','credit_card',2,7,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('audit_log','Audit Log','A','/audit-logs','audit_log','zoom_in',4,1,'admin');

insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('currency','Currency','A','/currencies','currency','local_atm',1,7,'setup');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,actions,parent) values ('locale','Locale','A','/locales','locale','public',1,7,'setup');

insert into roles (roleid, rolename, status, remark) values ('admin','Admin','A','Admin');
insert into roles (roleid, rolename, status, remark) values ('call_center','Call Center','A','Call Center');
insert into roles (roleid, rolename, status, remark) values ('it_support','IT Support','A','IT Support');
insert into roles (roleid, rolename, status, remark) values ('operator','Operator Group','A','Operator Group');

insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00001','gareth.bale','gareth.bale@gmail.com','Gareth Bale','https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00002','cristiano.ronaldo','cristiano.ronaldo@gmail.com','Cristiano Ronaldo','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00003','james.rodriguez','james.rodriguez@gmail.com','James Rodríguez','https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00004','zinedine.zidane','zinedine.zidane@gmail.com','Zinedine Zidane','https://upload.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00005','kaka','kaka@gmail.com','Kaká','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00006','luis.figo','luis.figo@gmail.com','Luís Figo','https://upload.wikimedia.org/wikipedia/commons/thumb/6/63/UEFA_TT_7209.jpg/440px-UEFA_TT_7209.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00007','ronaldo','ronaldo@gmail.com','Ronaldo','https://upload.wikimedia.org/wikipedia/commons/c/c8/Real_Valladolid-Valencia_CF%2C_2019-05-18_%2890%29_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00008','thibaut.courtois','thibaut.courtois@gmail.com','Thibaut Courtois','https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Courtois_2018_%28cropped%29.jpg/440px-Courtois_2018_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00009','luka.modric','luka.modric@gmail.com','Luka Modrić','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/ISL-HRV_%287%29.jpg/440px-ISL-HRV_%287%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00010','xabi.alonso','xabi.alonso@gmail.com','Xabi Alonso','https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg/440px-Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00011','karim.benzema','karim.benzema@gmail.com','Karim Benzema','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Karim_Benzema_2018.jpg/440px-Karim_Benzema_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00012','marc-andre.ter.stegen','marc-andre.ter.stegen@gmail.com','Marc-André ter Stegen','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e1/Marc-Andr%C3%A9_ter_Stegen.jpg/500px-Marc-Andr%C3%A9_ter_Stegen.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00013','sergino.dest','sergino.dest@gmail.com','Sergiño Dest','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Sergino_Dest.jpg/440px-Sergino_Dest.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00014','gerard.pique','gerard.pique@gmail.com','Gerard Piqué','https://upload.wikimedia.org/wikipedia/commons/4/4e/Gerard_Piqu%C3%A9_2018.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00015','ronald.araujo','ronald.araujo@gmail.com@gmail.com','Ronald Araújo','https://pbs.twimg.com/media/EtnqxaEU0AAc6A6.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00016','sergio.busquets','sergio.busquets@gmail.com@gmail.com','Sergio Busquets','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/Sergio_Busquets_2018.jpg/440px-Sergio_Busquets_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00017','antoine.griezmann','antoine.griezmann@gmail.com@gmail.com','Antoine Griezmann','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Antoine_Griezmann_in_2018_%28cropped%29.jpg/440px-Antoine_Griezmann_in_2018_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00018','miralem.pjanic','miralem.pjanic@gmail.com@gmail.com','Miralem Pjanić','https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg/440px-20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg','A','M','0987654321','Mrs','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00019','martin.braithwaite','martin.braithwaite@gmail.com@gmail.com','Martin Braithwaite','https://img.a.transfermarkt.technology/portrait/header/95732-1583334177.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00020','ousmane.dembele','ousmane.dembele@gmail.com@gmail.com','Ousmane Dembélé','https://upload.wikimedia.org/wikipedia/commons/7/77/Ousmane_Demb%C3%A9l%C3%A9_2018.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00021','riqui.puig','riqui.puig@gmail.com@gmail.com','Riqui Puig','https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Bar%C3%A7a_Napoli_12_%28cropped%29.jpg/440px-Bar%C3%A7a_Napoli_12_%28cropped%29.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00022','philip.coutinho','philip.coutinho@gmail.com@gmail.com','Philip Coutinho','https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Norberto_Murara_Neto_2019.jpg/440px-Norberto_Murara_Neto_2019.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00023','victor.lindelof','victor.lindelof@gmail.com@gmail.com','Victor Lindelöf','https://upload.wikimedia.org/wikipedia/commons/thumb/c/cc/CSKA-MU_2017_%286%29.jpg/440px-CSKA-MU_2017_%286%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00024','eric.bailly','eric.bailly@gmail.com@gmail.com','Eric Bailly','https://upload.wikimedia.org/wikipedia/commons/c/cf/Eric_Bailly_-_ManUtd.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00025','phil.jones','phil.jones@gmail.com@gmail.com','Phil Jones','https://upload.wikimedia.org/wikipedia/commons/f/fa/Phil_Jones_2018-06-28_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00026','harry.maguire','harry.maguire@gmail.com@gmail.com','Harry Maguire','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00027','paul.pogba','paul.pogba@gmail.com@gmail.com','Paul Pogba','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00028','edinson.cavani','edinson.cavani@gmail.com@gmail.com','Edinson Cavani','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Edinson_Cavani_2018.jpg/440px-Edinson_Cavani_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00029','juan.mata','juan.mata@gmail.com@gmail.com','Juan Mata','https://upload.wikimedia.org/wikipedia/commons/7/70/Ukr-Spain2015_%286%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00030','anthony.martial','anthony.martial@gmail.com@gmail.com','Anthony Martial','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Anthony_Martial_27_September_2017_cropped.jpg/440px-Anthony_Martial_27_September_2017_cropped.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00031','marcus.rashford','marcus.rashford@gmail.com@gmail.com','Marcus Rashford','https://upload.wikimedia.org/wikipedia/commons/5/5e/Press_Tren_CSKA_-_MU_%283%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00032','mason.greenwood','mason.greenwood@gmail.com@gmail.com','Mason Greenwood','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Mason_Greenwood.jpeg/440px-Mason_Greenwood.jpeg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00033','lee.grant','lee.grant@gmail.com@gmail.com','Lee Grant','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/LeeGrant09.jpg/400px-LeeGrant09.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00034','jesse.lingard','jesse.lingard@gmail.com@gmail.com','Jesse Lingard','https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Jesse_Lingard_2018-06-13_1.jpg/440px-Jesse_Lingard_2018-06-13_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00035','keylor.navas','keylor.navas@gmail.com@gmail.com','Keylor Navas','https://upload.wikimedia.org/wikipedia/commons/thumb/d/dc/Keylor_Navas_2018_%28cropped%29.jpg/220px-Keylor_Navas_2018_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00036','achraf.hakimi','achraf.hakimi@gmail.com@gmail.com','Achraf Hakimi','https://upload.wikimedia.org/wikipedia/commons/9/91/Iran-Morocco_by_soccer.ru_14_%28Achraf_Hakimi%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00037','presnel.kimpembe','presnel.kimpembe@gmail.com@gmail.com','Presnel Kimpembe','https://upload.wikimedia.org/wikipedia/commons/thumb/0/0e/Presnel_Kimpembe.jpg/400px-Presnel_Kimpembe.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00038','sergio.ramos','sergio.ramos@gmail.com@gmail.com','Sergio Ramos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg/440px-FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00039','marquinhos','marquinhos@gmail.com@gmail.com','Marquinhos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg/440px-Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00040','marco.verratti','marco.verratti@gmail.com@gmail.com','Marco Verratti','https://upload.wikimedia.org/wikipedia/commons/d/d0/Kiev-PSG_%289%29.jpg','A','M','0987654321','Mr','E');

update users set language = 'en', dateformat = 'd/M/yyyy';

insert into userroles(userid, roleid) values ('00001','admin');
insert into userroles(userid, roleid) values ('00003','admin');
insert into userroles(userid, roleid) values ('00004','admin');
insert into userroles(userid, roleid) values ('00005','it_support');
insert into userroles(userid, roleid) values ('00007','admin');
insert into userroles(userid, roleid) values ('00008','call_center');
insert into userroles(userid, roleid) values ('00009','it_support');
insert into userroles(userid, roleid) values ('00010','call_center');
insert into userroles(userid, roleid) values ('00011','it_support');
insert into userroles(userid, roleid) values ('00012','call_center');
insert into userroles(userid, roleid) values ('00012','it_support');

insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'setup', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'report', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'audit_log', 7);

insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'audit_log', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'setup', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'currency', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'locale', 7);

INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('6xydt3Qap', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-02 21:00:06.811', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gRAIVh1tM', 'term', '00005', '188.239.138.226', 'patch', '2023-07-03 12:09:51.659', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('d8sQRO1ap', 'entity', '00005', '188.239.138.226', 'patch', '2023-07-03 13:04:20.950', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gMu1Rh1aM', 'entity', '00005', '188.239.138.226', 'patch', '2023-07-03 13:04:24.491', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('jrFkzsQaM', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-03 16:00:42.627', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('4lVacs1aM', 'authentication', '00001', '::1', 'authenticate', '2023-07-03 16:22:13.157', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('a8Y-cbQtM', 'product', '00001', '95.194.49.166', 'patch', '2023-07-03 16:22:23.430', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Wvc4Us1aM', 'term', '00001', '95.194.49.166', 'patch', '2023-07-03 20:43:31.757', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('tcztIsQap', 'term', '00001', '::1', 'create', '2023-07-03 20:44:02.086', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('dO7zIb1ap', 'entity', '00001', '::1', 'patch', '2023-07-03 20:44:47.349', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('K-KcIbQtp', 'company', '00001', '::1', 'patch', '2023-07-03 20:45:55.702', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('G5JcIsQap', 'company', '00001', '::1', 'patch', '2023-07-03 20:45:59.129', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('HaLnIb1tM', 'company', '00001', '::1', 'patch', '2023-07-03 20:46:02.818', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('h_kcUbQap', 'company', '00001', '219.62.20.91', 'patch', '2023-07-03 20:46:05.519', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('jpTZIbQtM', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:07.779', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('UH_ZUsQtp', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:32.408', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('wP1SUsQtp', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:34.747', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('BxYPUb1aM', 'role', '00001', '::1', 'patch', '2023-07-03 20:46:42.944', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('rjegUs1tM', 'role', '00001', '::1', 'patch', '2023-07-03 20:47:02.120', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('lbmgUbQtM', 'role', '00001', '::1', 'patch', '2023-07-03 20:47:09.713', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('5o7-JsQap', 'role', '00001', '::1', 'patch', '2023-07-03 21:02:15.442', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('6eTFGbQap', 'role', '00001', '::1', 'patch', '2023-07-03 21:05:48.155', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('14S3JsQaM', 'role', '00001', '::1', 'patch', '2023-07-03 21:05:55.771', 'fail', 'pq: duplicate key text violates unique constraint "rolemodules_pkey"');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('DOYhJb1tp', 'article', '00001', '::1', 'patch', '2023-07-03 21:06:22.692', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gKzOGs1tp', 'article', '00001', '::1', 'patch', '2023-07-03 21:06:25.995', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('SD3OJsQaM', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-03 21:06:32.586', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('wD-7GbQaM', 'term', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:36.507', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('n3x7Js1tp', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:41.929', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Jm2NJbQap', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:47.577', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('mHJNJbQtM', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:54.878', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('u2RuJs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:09:32.212', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('2GrXJb1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:09:43.729', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('tx0dJsQtM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:10.950', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Ua9dJbQaM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:15.896', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('QD3KJb1tp', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:21.980', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('CU5dGs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:26.719', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('UnAKJs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:31.352', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('SiyKGs1ap', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:38.634', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('yYReJsQaM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:11:10.110', 'success', '');

insert into code_master(master, code, name, sequence, status) values ('language','en','English',1,'A');
insert into code_master(master, code, name, sequence, status) values ('language','vi','Tiếng Việt',2,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yyyy/M/d','yyyy/M/d',1,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yyyy/MM/dd','yyyy/MM/dd',2,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yyyy-M-d','yyyy-M-d',3,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yyyy-MM-dd','yyyy-MM-dd',4,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yyyy.MM.dd','yyyy.MM.dd',5,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','yy.MM.dd','yy.MM.dd',6,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','d/M/yyyy','d/M/yyyy',7,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','d/MM/yyyy','d/MM/yyyy',8,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd/MM/yyyy','dd/MM/yyyy',9,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd/MM yyyy','dd/MM yyyy',10,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd/MM/yy','dd/MM/yy',11,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','d-M-yyyy','d-M-yyyy',12,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd-MM-yyyy','dd-MM-yyyy',13,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd-MM-yy','dd-MM-yy',14,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','d.M.yyyy','d.M.yyyy',15,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','d.MM.yyyy','d.MM.yyyy',16,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd.MM.yyyy','dd.MM.yyyy',17,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','dd.MM.yy','dd.MM.yy',18,'I');
insert into code_master(master, code, name, sequence, status) values ('date_format','M/d/yyyy','M/d/yyyy',19,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','MM/dd/yyyy','MM/dd/yyyy',20,'A');
insert into code_master(master, code, name, sequence, status) values ('date_format','MM.dd.yyyy','MM.dd.yyyy',21,'A');

/*
alter table userroles add foreign key (userid) references users (userid);
alter table userroles add foreign key (roleid) references roles (roleid);

alter table modules add foreign key (parent) references modules (moduleid);

alter table rolemodules add foreign key (roleid) references roles (roleid);
alter table rolemodules add foreign key (moduleid) references modules (moduleid);

drop table modules;
drop table users;
drop table roles;
drop table userroles;
drop table rolemodules;
drop table auditlog;
*/

CREATE TABLE currency (
    code bpchar(3) NOT NULL PRIMARY KEY,
    symbol varchar(6) NOT NULL,
    decimal_digits int4 NULL,
    status char(1)
);
CREATE TABLE locale (
    code varchar(40) NOT NULL PRIMARY KEY,
    name varchar(255) NULL,
    native_name varchar(255) NULL,
    country_code varchar(5) NULL,
    country_name varchar(255) NULL,
    native_country_name varchar(255) NULL,
    date_format varchar(14) NULL,
    first_day_of_week int2 NULL,
    decimal_separator varchar(3) NULL,
    group_separator varchar(3) NULL,
    currency_code char(3) NULL,
    currency_symbol varchar(6) NULL,
    currency_decimal_digits int2 NULL,
    currency_pattern int2 NULL,
    currency_sample varchar(40) NULL
);

insert into currency(code,decimal_digits,symbol) values ('AED',2,'د.إ');
insert into currency(code,decimal_digits,symbol) values ('AFN',2,'؋');
insert into currency(code,decimal_digits,symbol) values ('ALL',2,'Lek');
insert into currency(code,decimal_digits,symbol) values ('AMD',2,'դր.');
insert into currency(code,decimal_digits,symbol) values ('ARS',2,'$');
insert into currency(code,decimal_digits,symbol) values ('AUD',2,'$');
insert into currency(code,decimal_digits,symbol) values ('AZN',2,'ман.');
insert into currency(code,decimal_digits,symbol) values ('BAM',2,'KM');
insert into currency(code,decimal_digits,symbol) values ('BDT',2,'৳');
insert into currency(code,decimal_digits,symbol) values ('BGN',2,'лв.');
insert into currency(code,decimal_digits,symbol) values ('BHD',3,'BD');
insert into currency(code,decimal_digits,symbol) values ('BND',0,'$');
insert into currency(code,decimal_digits,symbol) values ('BOB',2,'$b');
insert into currency(code,decimal_digits,symbol) values ('BRL',2,'R$');
insert into currency(code,decimal_digits,symbol) values ('BYR',2,'р.');
insert into currency(code,decimal_digits,symbol) values ('BZD',2,'BZ$');
insert into currency(code,decimal_digits,symbol) values ('CAD',2,'$');
insert into currency(code,decimal_digits,symbol) values ('CHF',2,'Fr.');
insert into currency(code,decimal_digits,symbol) values ('CLP',2,'$');
insert into currency(code,decimal_digits,symbol) values ('CNY',2,'¥');
insert into currency(code,decimal_digits,symbol) values ('COP',2,'$');
insert into currency(code,decimal_digits,symbol) values ('CRC',2,'₡');
insert into currency(code,decimal_digits,symbol) values ('CSD',2,'Дин.');
insert into currency(code,decimal_digits,symbol) values ('CZK',2,'Kč');
insert into currency(code,decimal_digits,symbol) values ('DKK',2,'kr.');
insert into currency(code,decimal_digits,symbol) values ('DOP',2,'RD$');
insert into currency(code,decimal_digits,symbol) values ('DZD',2,'DA');
insert into currency(code,decimal_digits,symbol) values ('EEK',2,'kr');
insert into currency(code,decimal_digits,symbol) values ('EGP',2,'£');
insert into currency(code,decimal_digits,symbol) values ('ETB',2,'Br');
insert into currency(code,decimal_digits,symbol) values ('EUR',2,'€');
insert into currency(code,decimal_digits,symbol) values ('GBP',2,'£');
insert into currency(code,decimal_digits,symbol) values ('GEL',2,'Lari');
insert into currency(code,decimal_digits,symbol) values ('GTQ',2,'Q');
insert into currency(code,decimal_digits,symbol) values ('HKD',2,'HK$');
insert into currency(code,decimal_digits,symbol) values ('HNL',2,'L.');
insert into currency(code,decimal_digits,symbol) values ('HRK',2,'kn');
insert into currency(code,decimal_digits,symbol) values ('HUF',2,'Ft');
insert into currency(code,decimal_digits,symbol) values ('IDR',0,'Rp');
insert into currency(code,decimal_digits,symbol) values ('ILS',2,'₪');
insert into currency(code,decimal_digits,symbol) values ('INR',2,'₹');
insert into currency(code,decimal_digits,symbol) values ('IQD',2,'ID');
insert into currency(code,decimal_digits,symbol) values ('IRR',2,'ريال');
insert into currency(code,decimal_digits,symbol) values ('ISK',0,'kr.');
insert into currency(code,decimal_digits,symbol) values ('JMD',2,'J$');
insert into currency(code,decimal_digits,symbol) values ('JOD',3,'د.أ');
insert into currency(code,decimal_digits,symbol) values ('JPY',0,'¥');
insert into currency(code,decimal_digits,symbol) values ('KES',2,'S');
insert into currency(code,decimal_digits,symbol) values ('KGS',2,'сом');
insert into currency(code,decimal_digits,symbol) values ('KHR',2,'៛');
insert into currency(code,decimal_digits,symbol) values ('KRW',0,'₩');
insert into currency(code,decimal_digits,symbol) values ('KWD',3,'KD');
insert into currency(code,decimal_digits,symbol) values ('KZT',2,'Т');
insert into currency(code,decimal_digits,symbol) values ('LAK',2,'₭');
insert into currency(code,decimal_digits,symbol) values ('LBP',2,'LL');
insert into currency(code,decimal_digits,symbol) values ('LKR',2,'රු.');
insert into currency(code,decimal_digits,symbol) values ('LTL',2,'Lt');
insert into currency(code,decimal_digits,symbol) values ('LVL',2,'Ls');
insert into currency(code,decimal_digits,symbol) values ('LYD',3,'LD');
insert into currency(code,decimal_digits,symbol) values ('MAD',2,'DH');
insert into currency(code,decimal_digits,symbol) values ('MKD',2,'ден.');
insert into currency(code,decimal_digits,symbol) values ('MNT',2,'₮');
insert into currency(code,decimal_digits,symbol) values ('MOP',2,'$');
insert into currency(code,decimal_digits,symbol) values ('MVR',2,'ރ.');
insert into currency(code,decimal_digits,symbol) values ('MXN',2,'$');
insert into currency(code,decimal_digits,symbol) values ('MYR',2,'RM');
insert into currency(code,decimal_digits,symbol) values ('NIO',2,'C$');
insert into currency(code,decimal_digits,symbol) values ('NOK',2,'kr');
insert into currency(code,decimal_digits,symbol) values ('NPR',2,'रु');
insert into currency(code,decimal_digits,symbol) values ('NZD',2,'$');
insert into currency(code,decimal_digits,symbol) values ('OMR',3,'R.O');
insert into currency(code,decimal_digits,symbol) values ('PAB',2,'B/.');
insert into currency(code,decimal_digits,symbol) values ('PEN',2,'S/.');
insert into currency(code,decimal_digits,symbol) values ('PHP',2,'₱');
insert into currency(code,decimal_digits,symbol) values ('PKR',2,'Rs');
insert into currency(code,decimal_digits,symbol) values ('PLN',2,'zł');
insert into currency(code,decimal_digits,symbol) values ('PYG',2,'Gs');
insert into currency(code,decimal_digits,symbol) values ('QAR',2,'QR');
insert into currency(code,decimal_digits,symbol) values ('RON',2,'lei');
insert into currency(code,decimal_digits,symbol) values ('RSD',2,'Дин.');
insert into currency(code,decimal_digits,symbol) values ('RUB',2,'һ.');
insert into currency(code,decimal_digits,symbol) values ('RWF',2,'R₣');
insert into currency(code,decimal_digits,symbol) values ('SAR',2,'SR');
insert into currency(code,decimal_digits,symbol) values ('SEK',2,'kr');
insert into currency(code,decimal_digits,symbol) values ('SGD',2,'$');
insert into currency(code,decimal_digits,symbol) values ('SYP',2,'LS');
insert into currency(code,decimal_digits,symbol) values ('THB',2,'฿');
insert into currency(code,decimal_digits,symbol) values ('TJS',2,'т.р.');
insert into currency(code,decimal_digits,symbol) values ('TMT',2,'m.');
insert into currency(code,decimal_digits,symbol) values ('TND',3,'DT');
insert into currency(code,decimal_digits,symbol) values ('TRY',2,'TL');
insert into currency(code,decimal_digits,symbol) values ('TTD',2,'TT$');
insert into currency(code,decimal_digits,symbol) values ('TWD',2,'NT$');
insert into currency(code,decimal_digits,symbol) values ('UAH',2,'₴');
insert into currency(code,decimal_digits,symbol) values ('USD',2,'$');
insert into currency(code,decimal_digits,symbol) values ('UYU',2,'$U');
insert into currency(code,decimal_digits,symbol) values ('UZS',0,'лв');
insert into currency(code,decimal_digits,symbol) values ('VEF',2,'Bs.');
insert into currency(code,decimal_digits,symbol) values ('VND',0,'₫');
insert into currency(code,decimal_digits,symbol) values ('XOF',2,'XOF');
insert into currency(code,decimal_digits,symbol) values ('YER',2,'﷼');
insert into currency(code,decimal_digits,symbol) values ('ZAR',2,'R');
insert into currency(code,decimal_digits,symbol) values ('ZWL',2,'Z$');

update currency set status = 'A';

INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('af-ZA','ZA','South Africa','Suid Afrika','Afrikaans (South Africa)','Afrikaans (Suid Afrika)','yyyy/MM/dd',1,'.',',',2,'ZAR',2,'R 10,000.00'),
('am-ET','ET','Ethiopia','ኢትዮጵያ','Amharic (Ethiopia)','አማርኛ (ኢትዮጵያ)','d/M/yyyy',1,'.',',',2,'ETB',0,'ETB10,000.00'),
('ar-AE','AE','U.A.E.','الإمارات العربية المتحدة','Arabic (U.A.E.)','العربية (الإمارات العربية المتحدة)','dd/MM/yyyy',7,'.',',',2,'AED',2,'د.إ.‏ 10,000.00'),
('ar-BH','BH','Bahrain','البحرين','Arabic (Bahrain)','العربية (البحرين)','dd/MM/yyyy',7,'.',',',3,'BHD',2,'د.ب.‏ 10,000.000'),
('ar-DZ','DZ','Algeria','الجزائر','Arabic (Algeria)','العربية (الجزائر)','dd-MM-yyyy',7,'.',',',2,'DZD',2,'د.ج.‏ 10,000.00'),
('ar-EG','EG','Egypt','مصر','Arabic (Egypt)','العربية (مصر)','dd/MM/yyyy',7,'.',',',2,'EGP',2,'ج.م.‏ 10,000.00'),
('ar-IQ','IQ','Iraq','العراق','Arabic (Iraq)','العربية (العراق)','dd/MM/yyyy',7,'.',',',2,'IQD',2,'د.ع.‏ 10,000.00'),
('ar-JO','JO','Jordan','الأردن','Arabic (Jordan)','العربية (الأردن)','dd/MM/yyyy',7,'.',',',3,'JOD',2,'د.ا.‏ 10,000.000'),
('ar-KW','KW','Kuwait','الكويت','Arabic (Kuwait)','العربية (الكويت)','dd/MM/yyyy',7,'.',',',3,'KWD',2,'د.ك.‏ 10,000.000'),
('ar-LB','LB','Lebanon','لبنان','Arabic (Lebanon)','العربية (لبنان)','dd/MM/yyyy',2,'.',',',2,'LBP',2,'ل.ل.‏ 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('ar-LY','LY','Libya','ليبيا','Arabic (Libya)','العربية (ليبيا)','dd/MM/yyyy',7,'.',',',3,'LYD',0,'د.ل.‏10,000.000'),
('ar-MA','MA','Morocco','المملكة المغربية','Arabic (Morocco)','العربية (المملكة المغربية)','dd-MM-yyyy',2,'.',',',2,'MAD',2,'د.م.‏ 10,000.00'),
('ar-OM','OM','Oman','عمان','Arabic (Oman)','العربية (عمان)','dd/MM/yyyy',7,'.',',',3,'OMR',2,'ر.ع.‏ 10,000.000'),
('ar-QA','QA','Qatar','قطر','Arabic (Qatar)','العربية (قطر)','dd/MM/yyyy',7,'.',',',2,'QAR',2,'ر.ق.‏ 10,000.00'),
('ar-SA','SA','Saudi Arabia','المملكة العربية السعودية','Arabic (Saudi Arabia)','العربية (المملكة العربية السعودية)','dd/MM/yy',7,'.',',',2,'SAR',2,'ر.س.‏ 10,000.00'),
('ar-SY','SY','Syria','سوريا','Arabic (Syria)','العربية (سوريا)','dd/MM/yyyy',7,'.',',',2,'SYP',2,'ل.س.‏ 10,000.00'),
('ar-TN','TN','Tunisia','تونس','Arabic (Tunisia)','العربية (تونس)','dd-MM-yyyy',2,'.',',',3,'TND',2,'د.ت.‏ 10,000.000'),
('ar-YE','YE','Yemen','اليمن','Arabic (Yemen)','العربية (اليمن)','dd/MM/yyyy',7,'.',',',2,'YER',2,'ر.ي.‏ 10,000.00'),
('arn-CL','CL','Chile','Chile','Mapudungun (Chile)','Mapudungun (Chile)','dd-MM-yyyy',1,',','.',2,'CLP',2,'$ 10.000,00'),
('as-IN','IN','India','ভাৰত','Assamese (India)','অসমীয়া (ভাৰত)','dd-MM-yyyy',2,'.',',',2,'INR',1,'10,000.00ট');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('az-Cyrl-AZ','AZ','Azerbaijan','Азәрбајҹан','Azeri (Cyrillic, Azerbaijan)','Азәрбајҹан (Азәрбајҹан)','dd.MM.yyyy',2,',',' ',2,'AZN',3,'10 000,00 ман.'),
('az-Latn-AZ','AZ','Azerbaijan','Azərbaycan','Azeri (Latin, Azerbaijan)','Azərbaycan­ılı (Azərbaycan)','dd.MM.yyyy',2,',',' ',2,'AZN',3,'10 000,00 man.'),
('ba-RU','RU','Russia','Россия','Bashkir (Russia)','Башҡорт (Россия)','dd.MM.yy',2,',',' ',2,'RUB',3,'10 000,00 һ.'),
('be-BY','BY','Belarus','Беларусь','Belarusian (Belarus)','Беларускі (Беларусь)','dd.MM.yyyy',2,',',' ',2,'BYR',3,'10 000,00 р.'),
('bg-BG','BG','Bulgaria','България','Bulgarian (Bulgaria)','български (България)','d.M.yyyy ''г.''',2,',',' ',2,'BGN',3,'10 000,00 лв.'),
('bn-BD','BD','Bangladesh','বাংলাদেশ','Bengali (Bangladesh)','বাংলা (বাংলাদেশ)','dd-MM-yy',2,'.',',',2,'BDT',2,'৳ 10,000.00'),
('bn-IN','IN','India','ভারত','Bengali (India)','বাংলা (ভারত)','dd-MM-yy',2,'.',',',2,'INR',2,'টা 10,000.00'),
('bo-CN','CN','People''s Republic of China','ཀྲུང་ཧྭ་མི་དམངས་སྤྱི་མཐུན་རྒྱལ་ཁབ།','Tibetan (PRC)','བོད་ཡིག (ཀྲུང་ཧྭ་མི་དམངས་སྤྱི་མཐུན་རྒྱལ་ཁབ།)','yyyy/M/d',2,'.',',',2,'CNY',0,'¥10,000.00'),
('br-FR','FR','France','Frañs','Breton (France)','brezhoneg (Frañs)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('bs-Cyrl-BA','BA','Bosnia and Herzegovina','Босна и Херцеговина','Bosnian (Cyrillic, Bosnia and Herzegovina)','босански (Босна и Херцеговина)','d.M.yyyy',2,',','.',2,'BAM',3,'10.000,00 КМ');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('bs-Latn-BA','BA','Bosnia and Herzegovina','Bosna i Hercegovina','Bosnian (Latin, Bosnia and Herzegovina)','bosanski (Bosna i Hercegovina)','d.M.yyyy',2,',','.',2,'BAM',3,'10.000,00 KM'),
('ca-ES','ES','Spain','Espanya','Catalan (Catalan)','català (català)','dd/MM/yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('co-FR','FR','France','France','Corsican (France)','Corsu (France)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('cs-CZ','CZ','Czech Republic','Česká republika','Czech (Czech Republic)','čeština (Česká republika)','d.M.yyyy',2,',',' ',2,'CZK',3,'10 000,00 Kč'),
('cy-GB','GB','United Kingdom','y Deyrnas Unedig','Welsh (United Kingdom)','Cymraeg (y Deyrnas Unedig)','dd/MM/yyyy',2,'.',',',2,'GBP',0,'£10,000.00'),
('da-DK','DK','Denmark','Danmark','Danish (Denmark)','dansk (Danmark)','dd-MM-yyyy',2,',','.',2,'DKK',2,'kr. 10.000,00'),
('de-AT','AT','Austria','Österreich','German (Austria)','Deutsch (Österreich)','dd.MM.yyyy',2,',','.',2,'EUR',2,'€ 10.000,00'),
('de-CH','CH','Switzerland','Schweiz','German (Switzerland)','Deutsch (Schweiz)','dd.MM.yyyy',2,'.','',2,'CHF',2,'Fr. 10''000.00'),
('de-DE','DE','Germany','Deutschland','German (Germany)','Deutsch (Deutschland)','dd.MM.yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('de-LI','LI','Liechtenstein','Liechtenstein','German (Liechtenstein)','Deutsch (Liechtenstein)','dd.MM.yyyy',2,'.','',2,'CHF',2,'CHF 10''000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('de-LU','LU','Luxembourg','Luxemburg','German (Luxembourg)','Deutsch (Luxemburg)','dd.MM.yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('dsb-DE','DE','Germany','Nimska','Lower Sorbian (Germany)','dolnoserbšćina (Nimska)','d. M. yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('dv-MV','MV','Maldives','ދިވެހި ރާއްޖެ','Divehi (Maldives)','ދިވެހިބަސް (ދިވެހި ރާއްޖެ)','dd/MM/yy',1,'.',',',2,'MVR',3,'10,000.00 ރ.'),
('el-GR','GR','Greece','Ελλάδα','Greek (Greece)','Ελληνικά (Ελλάδα)','d/M/yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('en-029','29','Caribbean','Caribbean','English (Caribbean)','English (Caribbean)','MM/dd/yyyy',2,'.',',',2,'USD',0,'$10,000.00'),
('en-AU','AU','Australia','Australia','English (Australia)','English (Australia)','d/MM/yyyy',2,'.',',',2,'AUD',0,'$10,000.00'),
('en-BZ','BZ','Belize','Belize','English (Belize)','English (Belize)','dd/MM/yyyy',1,'.',',',2,'BZD',0,'BZ$10,000.00'),
('en-CA','CA','Canada','Canada','English (Canada)','English (Canada)','dd/MM/yyyy',1,'.',',',2,'CAD',0,'$10,000.00'),
('en-GB','GB','United Kingdom','United Kingdom','English (United Kingdom)','English (United Kingdom)','dd/MM/yyyy',2,'.',',',2,'GBP',0,'£10,000.00'),
('en-IE','IE','Ireland','Ireland','English (Ireland)','English (Ireland)','dd/MM/yyyy',2,'.',',',2,'EUR',0,'€ 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('en-IN','IN','India','India','English (India)','English (India)','dd-MM-yyyy',2,'.',',',2,'INR',2,'Rs. 10,000.00'),
('en-JM','JM','Jamaica','Jamaica','English (Jamaica)','English (Jamaica)','dd/MM/yyyy',1,'.',',',2,'JMD',0,'J$10,000.00'),
('en-MY','MY','Malaysia','Malaysia','English (Malaysia)','English (Malaysia)','d/M/yyyy',1,'.',',',2,'MYR',0,'RM10,000.00'),
('en-NZ','NZ','New Zealand','New Zealand','English (New Zealand)','English (New Zealand)','d/MM/yyyy',2,'.',',',2,'NZD',0,'$10,000.00'),
('en-PH','PH','Philippines','Philippines','English (Republic of the Philippines)','English (Philippines)','M/d/yyyy',1,'.',',',2,'PHP',0,'Php10,000.00'),
('en-SG','SG','Singapore','Singapore','English (Singapore)','English (Singapore)','d/M/yyyy',1,'.',',',2,'SGD',0,'$10,000.00'),
('en-TT','TT','Trinidad and Tobago','Trinidad y Tobago','English (Trinidad and Tobago)','English (Trinidad y Tobago)','dd/MM/yyyy',1,'.',',',2,'TTD',0,'TT$10,000.00'),
('en-US','US','United States','United States','English (United States)','English (United States)','M/d/yyyy',1,'.',',',2,'USD',0,'$10,000.00'),
('en-ZA','ZA','South Africa','South Africa','English (South Africa)','English (South Africa)','yyyy/MM/dd',1,',',' ',2,'ZAR',2,'R 10 000.00'),
('en-ZW','ZW','Zimbabwe','Zimbabwe','English (Zimbabwe)','English (Zimbabwe)','M/d/yyyy',1,'.',',',2,'ZWL',0,'Z$10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('es-AR','AR','Argentina','Argentina','Spanish (Argentina)','Español (Argentina)','dd/MM/yyyy',1,',','.',2,'ARS',2,'$ 10.000,00'),
('es-BO','BO','Bolivia','Bolivia','Spanish (Bolivia)','Español (Bolivia)','dd/MM/yyyy',1,',','.',2,'BOB',2,'$b 10.000,00'),
('es-CL','CL','Chile','Chile','Spanish (Chile)','Español (Chile)','dd-MM-yyyy',1,',','.',2,'CLP',2,'$ 10.000,00'),
('es-CO','CO','Colombia','Colombia','Spanish (Colombia)','Español (Colombia)','dd/MM/yyyy',1,',','.',2,'COP',2,'$ 10.000,00'),
('es-CR','CR','Costa Rica','Costa Rica','Spanish (Costa Rica)','Español (Costa Rica)','dd/MM/yyyy',1,',','.',2,'CRC',0,'₡10.000,00'),
('es-DO','DO','Dominican Republic','República Dominicana','Spanish (Dominican Republic)','Español (República Dominicana)','dd/MM/yyyy',1,'.',',',2,'DOP',0,'RD$10,000.00'),
('es-EC','EC','Ecuador','Ecuador','Spanish (Ecuador)','Español (Ecuador)','dd/MM/yyyy',1,',','.',2,'USD',2,'$ 10.000,00'),
('es-ES','ES','Spain','España','Spanish (Spain, International Sort)','Español (España, alfabetización internacional)','dd/MM/yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('es-GT','GT','Guatemala','Guatemala','Spanish (Guatemala)','Español (Guatemala)','dd/MM/yyyy',1,'.',',',2,'GTQ',0,'Q10,000.00'),
('es-HN','HN','Honduras','Honduras','Spanish (Honduras)','Español (Honduras)','dd/MM/yyyy',1,'.',',',2,'HNL',2,'L. 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('es-MX','MX','Mexico','México','Spanish (Mexico)','Español (México)','dd/MM/yyyy',1,'.',',',2,'MXN',0,'$10,000.00'),
('es-NI','NI','Nicaragua','Nicaragua','Spanish (Nicaragua)','Español (Nicaragua)','dd/MM/yyyy',1,'.',',',2,'NIO',2,'C$ 10,000.00'),
('es-PA','PA','Panama','Panamá','Spanish (Panama)','Español (Panamá)','MM/dd/yyyy',1,'.',',',2,'PAB',2,'B/. 10,000.00'),
('es-PE','PE','Peru','Perú','Spanish (Peru)','Español (Perú)','dd/MM/yyyy',1,'.',',',2,'PEN',2,'S/. 10,000.00'),
('es-PR','PR','Puerto Rico','Puerto Rico','Spanish (Puerto Rico)','Español (Puerto Rico)','dd/MM/yyyy',1,'.',',',2,'USD',2,'$10,000.00'),
('es-PY','PY','Paraguay','Paraguay','Spanish (Paraguay)','Español (Paraguay)','dd/MM/yyyy',2,',','.',2,'PYG',2,'Gs 10.000,00'),
('es-SV','SV','El Salvador','El Salvador','Spanish (El Salvador)','Español (El Salvador)','dd/MM/yyyy',1,'.',',',2,'USD',0,'$10,000.00'),
('es-US','US','United States','Estados Unidos','Spanish (United States)','Español (Estados Unidos)','M/d/yyyy',1,'.',',',2,'USD',0,'$10,000.00'),
('es-UY','UY','Uruguay','Uruguay','Spanish (Uruguay)','Español (Uruguay)','dd/MM/yyyy',2,',','.',2,'UYU',2,'$U 10.000,00'),
('es-VE','VE','Bolivarian Republic of Venezuela','Republica Bolivariana de Venezuela','Spanish (Bolivarian Republic of Venezuela)','Español (Republica Bolivariana de Venezuela)','dd/MM/yyyy',1,',','.',2,'VEF',2,'Bs. F. 10.000,00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('et-EE','EE','Estonia','Eesti','Estonian (Estonia)','eesti (Eesti)','d.MM.yyyy',2,'.',' ',2,'EEK',3,'10 000,00 kr'),
('eu-ES','ES','Spain','Espainia','Basque (Basque)','euskara (euskara)','yyyy/MM/dd',2,',','.',2,'EUR',3,'10.000,00 €'),
('fa-IR','IR','Iran','ایران','Persian','فارسى (ایران)','MM/dd/yyyy',7,'/',',',2,'IRR',2,'ريال 10,000.00'),
('fi-FI','FI','Finland','Suomi','Finnish (Finland)','suomi (Suomi)','d.M.yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('fil-PH','PH','Philippines','Pilipinas','Filipino (Philippines)','Filipino (Pilipinas)','M/d/yyyy',1,'.',',',2,'PHP',0,'PhP10,000.00'),
('fo-FO','FO','Faroe Islands','Føroyar','Faroese (Faroe Islands)','føroyskt (Føroyar)','dd-MM-yyyy',2,',','.',2,'DKK',2,'kr. 10.000,00'),
('fr-BE','BE','Belgium','Belgique','French (Belgium)','français (Belgique)','d/MM/yyyy',2,',','.',2,'EUR',2,'€ 10.000,00'),
('fr-CA','CA','Canada','Canada','French (Canada)','français (Canada)','yyyy-MM-dd',1,',',' ',2,'CAD',3,'10 000,00 $'),
('fr-CH','CH','Switzerland','Suisse','French (Switzerland)','français (Suisse)','dd.MM.yyyy',2,'.','',2,'CHF',2,'fr. 10''000.00'),
('fr-FR','FR','France','France','French (France)','français (France)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('fr-LU','LU','Luxembourg','Luxembourg','French (Luxembourg)','français (Luxembourg)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('fr-MC','MC','Principality of Monaco','Principauté de Monaco','French (Monaco)','français (Principauté de Monaco)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('fy-NL','NL','Netherlands','Nederlân','Frisian (Netherlands)','Frysk (Nederlân)','d-M-yyyy',2,',','.',2,'EUR',2,'€ 10.000,00'),
('ga-IE','IE','Ireland','Éire','Irish (Ireland)','Gaeilge (Éire)','dd/MM/yyyy',2,'.',',',2,'EUR',0,'€ 10,000.00'),
('gd-GB','GB','United Kingdom','An Rìoghachd Aonaichte','Scottish Gaelic (United Kingdom)','Gàidhlig (An Rìoghachd Aonaichte)','dd/MM/yyyy',2,'.',',',2,'GBP',0,'£10,000.00'),
('gl-ES','ES','Spain','España','Galician (Galician)','galego (galego)','dd/MM/yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('gsw-FR','FR','France','Frànkrisch','Alsatian (France)','Elsässisch (Frànkrisch)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('gu-IN','IN','India','ભારત','Gujarati (India)','ગુજરાતી (ભારત)','dd-MM-yy',2,'.',',',2,'INR',2,'રૂ 10,000.00'),
('ha-Latn-NG','NG','Nigeria','Nigeria','Hausa (Latin, Nigeria)','Hausa (Nigeria)','d/M/yyyy',1,'.',',',2,'NIO',2,'N 10,000.00'),
('he-IL','IL','Israel','ישראל','Hebrew (Israel)','עברית (ישראל)','dd/MM/yyyy',1,'.',',',2,'ILS',2,'₪ 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('hi-IN','IN','India','भारत','Hindi (India)','हिंदी (भारत)','dd-MM-yyyy',2,'.',',',2,'INR',2,'रु 10,000.00'),
('hr-BA','BA','Bosnia and Herzegovina','Bosna i Hercegovina','Croatian (Latin, Bosnia and Herzegovina)','hrvatski (Bosna i Hercegovina)','d.M.yyyy.',2,',','.',2,'BAM',3,'10.000,00 KM'),
('hr-HR','HR','Croatia','Hrvatska','Croatian (Croatia)','hrvatski (Hrvatska)','d.M.yyyy.',2,',','.',2,'HRK',3,'10.000,00 kn'),
('hsb-DE','DE','Germany','Němska','Upper Sorbian (Germany)','hornjoserbšćina (Němska)','d. M. yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('hu-HU','HU','Hungary','Magyarország','Hungarian (Hungary)','magyar (Magyarország)','yyyy.MM.dd.',2,',',' ',2,'HUF',3,'10 000,00 Ft'),
('hy-AM','AM','Armenia','Հայաստան','Armenian (Armenia)','Հայերեն (Հայաստան)','dd.MM.yyyy',2,'.',',',2,'AMD',3,'10,000.00 դր.'),
('id-ID','ID','Indonesia','Indonesia','Indonesian (Indonesia)','Bahasa Indonesia (Indonesia)','dd/MM/yyyy',2,',','.',0,'IDR',0,'Rp10.000'),
('ig-NG','NG','Nigeria','Nigeria','Igbo (Nigeria)','Igbo (Nigeria)','d/M/yyyy',1,'.',',',2,'NIO',2,'N 10,000.00'),
('ii-CN','CN','People''s Republic of China','ꍏꉸꏓꂱꇭꉼꇩ','Yi (PRC)','ꆈꌠꁱꂷ (ꍏꉸꏓꂱꇭꉼꇩ)','yyyy/M/d',2,'.',',',2,'CNY',0,'¥10,000.00'),
('is-IS','IS','Iceland','Ísland','Icelandic (Iceland)','íslenska (Ísland)','d.M.yyyy',2,',','.',0,'ISK',3,'10.000 kr.');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('it-CH','CH','Switzerland','Svizzera','Italian (Switzerland)','italiano (Svizzera)','dd.MM.yyyy',2,'.','',2,'CHF',2,'fr. 10''000.00'),
('it-IT','IT','Italy','Italia','Italian (Italy)','italiano (Italia)','dd/MM/yyyy',2,',','.',2,'EUR',2,'€ 10.000,00'),
('iu-Cans-CA','CA','Canada','ᑲᓇᑕ','Inuktitut (Syllabics, Canada)','ᐃᓄᒃᑎᑐᑦ (ᑲᓇᑕᒥ)','d/M/yyyy',1,'.',',',2,'CAD',0,'$10,000.00'),
('iu-Latn-CA','CA','Canada','kanata','Inuktitut (Latin, Canada)','Inuktitut (Kanatami)','d/MM/yyyy',1,'.',',',2,'CAD',0,'$10,000.00'),
('ja-JP','JP','Japan','日本','Japanese (Japan)','日本語 (日本)','yyyy/MM/dd',1,'.',',',0,'JPY',0,'¥10,000'),
('ka-GE','GE','Georgia','საქართველო','Georgian (Georgia)','ქართული (საქართველო)','dd.MM.yyyy',2,',',' ',2,'GEL',3,'10 000,00 Lari'),
('kk-KZ','KZ','Kazakhstan','Қазақстан','Kazakh (Kazakhstan)','Қазақ (Қазақстан)','dd.MM.yyyy',2,'-',' ',2,'KZT',0,'Т10 000,00'),
('kl-GL','GL','Greenland','Kalaallit Nunaat','Greenlandic (Greenland)','kalaallisut (Kalaallit Nunaat)','dd-MM-yyyy',2,',','.',2,'DKK',2,'kr. 10.000,00'),
('km-KH','KH','Cambodia','កម្ពុជា','Khmer (Cambodia)','ខ្មែរ (កម្ពុជា)','yyyy-MM-dd',1,'.',',',2,'KHR',1,'10,000.00៛'),
('kn-IN','IN','India','ಭಾರತ','Kannada (India)','ಕನ್ನಡ (ಭಾರತ)','dd-MM-yy',2,'.',',',2,'INR',2,'ರೂ 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('ko-KR','KR','Korea','대한민국','Korean (Korea)','한국어 (대한민국)','yyyy-MM-dd',1,'.',',',0,'KRW',0,'₩10,000'),
('kok-IN','IN','India','भारत','Konkani (India)','कोंकणी (भारत)','dd-MM-yyyy',2,'.',',',2,'INR',2,'रु 10,000.00'),
('ky-KG','KG','Kyrgyzstan','Кыргызстан','Kyrgyz (Kyrgyzstan)','Кыргыз (Кыргызстан)','dd.MM.yy',2,'-',' ',2,'KGS',3,'10 000,00 сом'),
('lb-LU','LU','Luxembourg','Luxembourg','Luxembourgish (Luxembourg)','Lëtzebuergesch (Luxembourg)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('lo-LA','LA','Lao P.D.R.','ສ.ປ.ປ. ລາວ','Lao (Lao P.D.R.)','ລາວ (ສ.ປ.ປ. ລາວ)','dd/MM/yyyy',1,'.',',',2,'LAK',1,'10,000.00₭'),
('lt-LT','LT','Lithuania','Lietuva','Lithuanian (Lithuania)','lietuvių (Lietuva)','yyyy.MM.dd',2,',','.',2,'LTL',3,'10.000,00 Lt'),
('lv-LV','LV','Latvia','Latvija','Latvian (Latvia)','latviešu (Latvija)','yyyy.MM.dd.',2,',',' ',2,'LVL',2,'Ls 10 000,00'),
('mi-NZ','NZ','New Zealand','Aotearoa','Maori (New Zealand)','Reo Māori (Aotearoa)','dd/MM/yyyy',2,'.',',',2,'NZD',0,'$10,000.00'),
('mk-MK','MK','Macedonia (FYROM)','Македонија','Macedonian (Former Yugoslav Republic of Macedonia)','македонски јазик (Македонија)','dd.MM.yyyy',2,',','.',2,'MKD',3,'10.000,00 ден.'),
('ml-IN','IN','India','ഭാരതം','Malayalam (India)','മലയാളം (ഭാരതം)','dd-MM-yy',2,'.',',',2,'INR',2,'ക 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('mn-MN','MN','Mongolia','Монгол улс','Mongolian (Cyrillic, Mongolia)','Монгол хэл (Монгол улс)','yy.MM.dd',2,',',' ',2,'MNT',1,'10 000,00₮'),
('mn-Mong-CN','CN','People''s Republic of China','ᠪᠦᠭᠦᠳᠡ ᠨᠠᠢᠷᠠᠮᠳᠠᠬᠤ ᠳᠤᠮᠳᠠᠳᠤ ᠠᠷᠠᠳ ᠣᠯᠣᠰ','Mongolian (Traditional Mongolian, PRC)','ᠮᠤᠨᠭᠭᠤᠯ ᠬᠡᠯᠡ (ᠪᠦᠭᠦᠳᠡ ᠨᠠᠢᠷᠠᠮᠳᠠᠬᠤ ᠳᠤᠮᠳᠠᠳᠤ ᠠᠷᠠᠳ ᠣᠯᠣᠰ)','yyyy/M/d',2,'.',',',2,'CNY',0,'¥10,000.00'),
('moh-CA','CA','Canada','Canada','Mohawk (Mohawk)','Kanien''kéha','M/d/yyyy',1,'.',',',2,'CAD',0,'$10,000.00'),
('mr-IN','IN','India','भारत','Marathi (India)','मराठी (भारत)','dd-MM-yyyy',2,'.',',',2,'INR',2,'रु 10,000.00'),
('ms-BN','BN','Brunei Darussalam','Brunei Darussalam','Malay (Brunei Darussalam)','Bahasa Melayu (Brunei Darussalam)','dd/MM/yyyy',2,',','.',0,'BND',0,'$10.00'),
('ms-MY','MY','Malaysia','Malaysia','Malay (Malaysia)','Bahasa Melayu (Malaysia)','dd/MM/yyyy',2,'.',',',0,'MYR',0,'RM10,000'),
('mt-MT','MT','Malta','Malta','Maltese (Malta)','Malti (Malta)','dd/MM/yyyy',2,'.',',',2,'EUR',0,'€ 10,000.00'),
('nb-NO','NO','Norway','Norge','Norwegian, Bokmål (Norway)','norsk, bokmål (Norge)','dd.MM.yyyy',2,',',' ',2,'NOK',2,'kr 10 000,00'),
('ne-NP','NP','Nepal','नेपाल','Nepali (Nepal)','नेपाली (नेपाल)','M/d/yyyy',1,'.',',',2,'NPR',0,'रु10,000.00'),
('nl-BE','BE','Belgium','België','Dutch (Belgium)','Nederlands (België)','d/MM/yyyy',2,',','.',2,'EUR',2,'€ 10.000,00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('nl-NL','NL','Netherlands','Nederland','Dutch (Netherlands)','Nederlands (Nederland)','d-M-yyyy',2,',','.',2,'EUR',2,'€ 10.000,00'),
('nn-NO','NO','Norway','Noreg','Norwegian, Nynorsk (Norway)','norsk, nynorsk (Noreg)','dd.MM.yyyy',2,',',' ',2,'NOK',2,'kr 10 000,00'),
('nso-ZA','ZA','South Africa','Afrika Borwa','Sesotho sa Leboa (South Africa)','Sesotho sa Leboa (Afrika Borwa)','yyyy/MM/dd',1,'.',',',2,'ZAR',2,'R 10,000.00'),
('oc-FR','FR','France','França','Occitan (France)','Occitan (França)','dd/MM/yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('or-IN','IN','India','ଭାରତ','Oriya (India)','ଓଡ଼ିଆ (ଭାରତ)','dd-MM-yy',1,'.',',',2,'INR',2,'ଟ 10,000.00'),
('pa-IN','IN','India','ਭਾਰਤ','Punjabi (India)','ਪੰਜਾਬੀ (ਭਾਰਤ)','dd-MM-yy',2,'.',',',2,'INR',2,'ਰੁ 10,000.00'),
('pl-PL','PL','Poland','Polska','Polish (Poland)','polski (Polska)','yyyy-MM-dd',2,',',' ',2,'PLN',3,'10 000,00 zł'),
('prs-AF','AF','Afghanistan','افغانستان','Dari (Afghanistan)','درى (افغانستان)','dd/MM/yy',6,'.',',',2,'AFN',0,'؋10.000,00'),
('ps-AF','AF','Afghanistan','افغانستان','Pashto (Afghanistan)','پښتو (افغانستان)','dd/MM/yy',7,'٫','٬',2,'AFN',0,'؋10،000,00'),
('pt-BR','BR','Brazil','Brasil','Portuguese (Brazil)','Português (Brasil)','dd/MM/yyyy',1,',','.',2,'BRL',2,'R$ 10.000,00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('pt-PT','PT','Portugal','Portugal','Portuguese (Portugal)','português (Portugal)','dd-MM-yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('qut-GT','GT','Guatemala','Guatemala','K''iche (Guatemala)','K''iche (Guatemala)','dd/MM/yyyy',1,'.',',',2,'GTQ',0,'Q10,000.00'),
('quz-BO','BO','Bolivia','Bolivia Suyu','Quechua (Bolivia)','runasimi (Qullasuyu)','dd/MM/yyyy',1,',','.',2,'BOB',2,'$b 10.000,00'),
('quz-EC','EC','Ecuador','Ecuador Suyu','Quechua (Ecuador)','runasimi (Ecuador)','dd/MM/yyyy',1,',','.',2,'USD',2,'$ 10.000,00'),
('quz-PE','PE','Peru','Peru Suyu','Quechua (Peru)','runasimi (Piruw)','dd/MM/yyyy',1,'.',',',2,'PEN',2,'S/. 10,000.00'),
('rm-CH','CH','Switzerland','Svizra','Romansh (Switzerland)','Rumantsch (Svizra)','dd/MM/yyyy',2,'.','',2,'CHF',2,'fr. 10''000.00'),
('ro-RO','RO','Romania','România','Romanian (Romania)','română (România)','dd.MM.yyyy',2,',','.',2,'RON',3,'10.000,00 lei'),
('ru-RU','RU','Russia','Россия','Russian (Russia)','русский (Россия)','dd.MM.yyyy',2,',',' ',2,'RUB',1,'10 000,00р.'),
('rw-RW','RW','Rwanda','Rwanda','Kinyarwanda (Rwanda)','Kinyarwanda (Rwanda)','M/d/yyyy',1,',',' ',2,'RWF',2,'RWF 10 000,00'),
('sa-IN','IN','India','भारतम्','Sanskrit (India)','संस्कृत (भारतम्)','dd-MM-yyyy',1,'.',',',2,'INR',2,'रु 10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('sah-RU','RU','Russia','Россия','Yakut (Russia)','саха (Россия)','MM.dd.yyyy',2,',',' ',2,'RUB',1,'10 000,00с.'),
('se-FI','FI','Finland','Suopma','Sami, Northern (Finland)','davvisámegiella (Suopma)','d.M.yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('se-NO','NO','Norway','Norga','Sami, Northern (Norway)','davvisámegiella (Norga)','dd.MM.yyyy',2,',',' ',2,'NOK',2,'kr 10 000,00'),
('se-SE','SE','Sweden','Ruoŧŧa','Sami, Northern (Sweden)','davvisámegiella (Ruoŧŧa)','yyyy-MM-dd',2,',','.',2,'SEK',3,'10 000,00 kr'),
('si-LK','LK','Sri Lanka','ශ්‍රී ලංකා','Sinhala (Sri Lanka)','සිංහල (ශ්‍රී ලංකා)','yyyy-MM-dd',2,'.',',',2,'LKR',2,'රු. 10,000.00'),
('sk-SK','SK','Slovakia','Slovenská republika','Slovak (Slovakia)','slovenčina (Slovenská republika)','d. M. yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('sl-SI','SI','Slovenia','Slovenija','Slovenian (Slovenia)','slovenski (Slovenija)','d.M.yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('sma-NO','NO','Norway','Nöörje','Sami, Southern (Norway)','åarjelsaemiengiele (Nöörje)','dd.MM.yyyy',2,',',' ',2,'NOK',2,'kr 10 000,00'),
('sma-SE','SE','Sweden','Sveerje','Sami, Southern (Sweden)','åarjelsaemiengiele (Sveerje)','yyyy-MM-dd',2,',','.',2,'SEK',3,'10 000,00 kr'),
('smj-NO','NO','Norway','Vuodna','Sami, Lule (Norway)','julevusámegiella (Vuodna)','dd.MM.yyyy',2,',',' ',2,'NOK',2,'kr 10 000,00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('smj-SE','SE','Sweden','Svierik','Sami, Lule (Sweden)','julevusámegiella (Svierik)','yyyy-MM-dd',2,',','.',2,'SEK',3,'10 000,00 kr'),
('smn-FI','FI','Finland','Suomâ','Sami, Inari (Finland)','sämikielâ (Suomâ)','d.M.yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('sms-FI','FI','Finland','Lää´ddjânnam','Sami, Skolt (Finland)','sääm´ǩiõll (Lää´ddjânnam)','d.M.yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('sq-AL','AL','Albania','Shqipëria','Albanian (Albania)','shqipe (Shqipëria)','yyyy-MM-dd',2,',','.',2,'ALL',1,'10.000,00Lek'),
('sr-Cyrl-BA','BA','Bosnia and Herzegovina','Босна и Херцеговина','Serbian (Cyrillic, Bosnia and Herzegovina)','српски (Босна и Херцеговина)','d.M.yyyy',2,',','.',2,'BAM',3,'10.000,00 КМ'),
('sr-Cyrl-CS','CS','Serbia and Montenegro (Former)','Србија и Црна Гора (Претходно)','Serbian (Cyrillic, Serbia and Montenegro (Former))','српски (Србија и Црна Гора (Претходно))','d.M.yyyy',2,',','.',2,'CSD',3,'10.000,00 Дин.'),
('sr-Cyrl-ME','ME','Montenegro','Црна Гора','Serbian (Cyrillic, Montenegro)','српски (Црна Гора)','d.M.yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('sr-Cyrl-RS','RS','Serbia','Србија','Serbian (Cyrillic, Serbia)','српски (Србија)','d.M.yyyy',2,',','.',2,'RSD',3,'10.000,00 Дин.'),
('sr-Latn-BA','BA','Bosnia and Herzegovina','Bosna i Hercegovina','Serbian (Latin, Bosnia and Herzegovina)','srpski (Bosna i Hercegovina)','d.M.yyyy',2,',','.',2,'BAM',3,'10.000,00 KM'),
('sr-Latn-CS','CS','Serbia and Montenegro (Former)','Srbija i Crna Gora (Prethodno)','Serbian (Latin, Serbia and Montenegro (Former))','srpski (Srbija i Crna Gora (Prethodno))','d.M.yyyy',2,',','.',2,'CSD',3,'10.000,00 Din.');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('sr-Latn-ME','ME','Montenegro','Crna Gora','Serbian (Latin, Montenegro)','srpski (Crna Gora)','d.M.yyyy',2,',','.',2,'EUR',3,'10.000,00 €'),
('sr-Latn-RS','RS','Serbia','Srbija','Serbian (Latin, Serbia)','srpski (Srbija)','d.M.yyyy',2,',','.',2,'RSD',3,'10.000,00 Din.'),
('sv-FI','FI','Finland','Finland','Swedish (Finland)','svenska (Finland)','d.M.yyyy',2,',',' ',2,'EUR',3,'10 000,00 €'),
('sv-SE','SE','Sweden','Sverige','Swedish (Sweden)','svenska (Sverige)','yyyy-MM-dd',2,',','.',2,'SEK',3,'10 000,00 kr'),
('sw-KE','KE','Kenya','Kenya','Kiswahili (Kenya)','Kiswahili (Kenya)','M/d/yyyy',1,'.',',',2,'KES',0,'S10,000.00'),
('syr-SY','SY','Syria','سوريا','Syriac (Syria)','ܣܘܪܝܝܐ (سوريا)','dd/MM/yyyy',7,'.',',',2,'SYP',2,'ل.س.‏ 10,000.00'),
('ta-IN','IN','India','இந்தியா','Tamil (India)','தமிழ் (இந்தியா)','dd-MM-yyyy',2,'.',',',2,'INR',2,'ரூ 10,000.00'),
('te-IN','IN','India','భారత దేశం','Telugu (India)','తెలుగు (భారత దేశం)','dd-MM-yy',2,'.',',',2,'INR',2,'రూ 10,000.00'),
('tg-Cyrl-TJ','TJ','Tajikistan','Тоҷикистон','Tajik (Cyrillic, Tajikistan)','Тоҷикӣ (Тоҷикистон)','dd.MM.yy',1,';',' ',2,'TJS',3,'10 000,00 т.р.'),
('th-TH','TH','Thailand','ไทย','Thai (Thailand)','ไทย (ไทย)','d/M/yyyy',2,'.',',',2,'THB',0,'฿10,000.00');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('tk-TM','TM','Turkmenistan','Türkmenistan','Turkmen (Turkmenistan)','türkmençe (Türkmenistan)','dd.MM.yy',2,',',' ',2,'TMT',1,'10 000,00m.'),
('tn-ZA','ZA','South Africa','Aforika Borwa','Setswana (South Africa)','Setswana (Aforika Borwa)','yyyy/MM/dd',1,'.',',',2,'ZAR',2,'R 10,000.00'),
('tr-TR','TR','Turkey','Türkiye','Turkish (Turkey)','Türkçe (Türkiye)','dd.MM.yyyy',2,',','.',2,'TRY',3,'10.000,00 TL'),
('tt-RU','RU','Russia','Россия','Tatar (Russia)','Татар (Россия)','dd.MM.yyyy',2,',',' ',2,'RUB',3,'10 000,00 р.'),
('tzm-Latn-DZ','DZ','Algeria','Djazaïr','Tamazight (Latin, Algeria)','Tamazight (Djazaïr)','dd-MM-yyyy',7,'.',',',2,'DZD',3,'10.000,00 DZD'),
('ug-CN','CN','People''s Republic of China','جۇڭخۇا خەلق جۇمھۇرىيىتى','Uyghur (PRC)','ئۇيغۇرچە (جۇڭخۇا خەلق جۇمھۇرىيىتى)','yyyy-M-d',1,'.',',',2,'CNY',0,'¥10,000.00'),
('uk-UA','UA','Ukraine','Україна','Ukrainian (Ukraine)','українська (Україна)','dd.MM.yyyy',2,',',' ',2,'UAH',1,'10 000,00₴'),
('ur-PK','PK','Islamic Republic of Pakistan','پاکستان','Urdu (Islamic Republic of Pakistan)','اُردو (پاکستان)','dd/MM/yyyy',2,'.',',',2,'PKR',0,'Rs10,000.00'),
('uz-Cyrl-UZ','UZ','Uzbekistan','Ўзбекистон Республикаси','Uzbek (Cyrillic, Uzbekistan)','Ўзбек (Ўзбекистон)','dd.MM.yyyy',2,',',' ',2,'UZS',3,'10 000,00 сўм'),
('uz-Latn-UZ','UZ','Uzbekistan','U''zbekiston Respublikasi','Uzbek (Latin, Uzbekistan)','U''zbek (U''zbekiston Respublikasi)','dd/MM yyyy',2,',',' ',0,'UZS',3,'10 000 so''m');
INSERT INTO locale (code,country_code,country_name,native_country_name,"name",native_name,date_format,first_day_of_week,decimal_separator,group_separator,currency_decimal_digits,currency_code,currency_pattern,currency_sample) VALUES
('vi-VN','VN','Vietnam','Việt Nam','Vietnamese (Vietnam)','Tiếng Việt (Việt Nam)','dd/MM/yyyy',2,',','.',2,'VND',3,'10.000,00 ₫'),
('wo-SN','SN','Senegal','Sénégal','Wolof (Senegal)','Wolof (Sénégal)','dd/MM/yyyy',2,',',' ',2,'XOF',3,'10 000,00 XOF'),
('xh-ZA','ZA','South Africa','uMzantsi Afrika','isiXhosa (South Africa)','isiXhosa (uMzantsi Afrika)','yyyy/MM/dd',1,'.',',',2,'ZAR',2,'R 10,000.00'),
('yo-NG','NG','Nigeria','Nigeria','Yoruba (Nigeria)','Yoruba (Nigeria)','d/M/yyyy',1,'.',',',2,'NIO',2,'N 10,000.00'),
('zh-CN','CN','People''s Republic of China','中华人民共和国','Chinese (Simplified, PRC)','中文(中华人民共和国)','yyyy/M/d',1,'.',',',2,'CNY',0,'¥10,000.00'),
('zh-HK','HK','Hong Kong S.A.R.','香港特別行政區','Chinese (Traditional, Hong Kong S.A.R.)','中文(香港特別行政區)','d/M/yyyy',1,'.',',',2,'HKD',0,'HK$10,000.00'),
('zh-MO','MO','Macao S.A.R.','澳門特別行政區','Chinese (Traditional, Macao S.A.R.)','中文(澳門特別行政區)','d/M/yyyy',1,'.',',',2,'MOP',0,'MOP10,000.00'),
('zh-SG','SG','Singapore','新加坡','Chinese (Simplified, Singapore)','中文(新加坡)','d/M/yyyy',1,'.',',',2,'SGD',0,'$10,000.00'),
('zh-TW','TW','Taiwan','台灣','Chinese (Traditional, Taiwan)','中文(台灣)','yyyy/M/d',1,'.',',',2,'TWD',0,'NT$10,000.00'),
('zu-ZA','ZA','South Africa','iNingizimu Afrika','isiZulu (South Africa)','isiZulu (iNingizimu Afrika)','yyyy/MM/dd',1,'.',',',2,'ZAR',2,'R 10,000.00');

UPDATE locale l
SET currency_symbol = c.symbol
    FROM currency c
where l.currency_code = c.code;