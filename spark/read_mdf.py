from asammdf import MDF, Signal
import mdfreader

mdf = MDF('/home/ec2-user/data/1CFD_VDDM_Lmbd1045_3.100_310_MS_log_SREC_log_shot0001_20190928_142405.dat')
metadata = mdf.info()

print('Groups found in file:')
for group in metadata.keys():
    print(group)
    is_dict = False
    try:
        metadata[group].keys()
        is_dict = True
    except:
        pass
    if is_dict:
#        print(metadata[group]['channel 1'])
        for i in range(1, metadata[group]['channels count']):
            print(metadata[group]['channel '+str(i)])
    print('='*80)
