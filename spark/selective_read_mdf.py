from asammdf import MDF, Signal
import mdfreader
import os
import pandas as pd
from functools import reduce
from pandas.api.types import is_numeric_dtype
from sklearn.preprocessing import OneHotEncoder


DATAPATH='/home/ec2-user/data/'
FILE='1CFD_VDDM_Lmbd1045_3.100_310_MS_log_SREC_log_shot0001_20190928_142405.dat'

def has_datadir(filename):
    """Check if filename contains path to datadir."""

    if DATAPATH in filename:
        return True
    return False

def add_datadir(filename):
    """Add path to datadir if not already in filename."""

    if not has_datadir(filename):
        return os.path.join(DATAPATH, FILE)
    else:
        return filename

def groups_in_file(mdf):
    """Return the number of groups in the MDF instance provided."""

    metadata = mdf.info()
    data_groups = 0
    for group in metadata.keys():
        is_dict = False
        try:
            metadata[group].keys()
            is_dict = True
        except:
            pass
        if is_dict:
            data_groups += 1
    return data_groups

def display_metadata(mdf, signals_only=False):
    """Display the metadata of the MDF instance provided."""
    print('Groups:')
    for group in metadata.keys():
        print(group)
        is_dict = False
        try:
            metadata[group].keys()
            is_dict = True
        except:
            pass
        if is_dict:
            if signals_only:
                for i in range(1, metadata[group]['channels count']):
                    print(metadata[group]['channel '+str(i)])
            else:
                print(metadata[group])
        print('='*80)

mdf = MDF(add_datadir(FILE))
#display_metadata(mdf, signals_only=True)
data_groups = groups_in_file(mdf)
print(data_groups, 'were found!')

#mdf = MDF(add_datadir(FILE), channels=required_channels)
#display_metadata(mdf)
#for g in range(data_groups):

dataframes = []
non_numeric_dataframes = []
for g in range(data_groups):
    df = mdf.get_group(g)
    if is_numeric_dtype(df.loc[0]):
        dataframes.append(df)
#        dataframes.append(df.iloc[0:10, :])
        print(dataframes[-1])
    else:
        non_numeric_dataframes.append(df)
print(len(dataframes), "numerical dataframes found!")
print(len(non_numeric_dataframes), "non-numerical dataframes found!")

def hot_encode(input_df):
    """Hot encode the given dataframe."""
    encoder = OneHotEncoder()
    encoder_df = pd.DataFrame(encoder.fit_transform(input_df).toarray(), index=input_df.index)
    column_names = []
    for i in range(encoder_df.shape[1]):
        column_names.append(input_df.columns[0])
    encoder_df.columns = column_names
    return encoder_df

#encoded_df = hot_encode(dataframes[3])

df_merged = reduce(lambda  left,right: pd.merge(left, right, on=['timestamps'],
    how='outer'), dataframes)
print("Result of merging", df_merged, df_merged.shape)


df_merged.to_csv('out.csv',index=True)
